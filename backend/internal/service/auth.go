package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/DangDDT/tada-learn-english/backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailExists        = errors.New("email already registered")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidToken       = errors.New("invalid refresh token")
	ErrTokenExpired       = errors.New("reset token expired or already used")
)

type AuthService struct {
	pool             *pgxpool.Pool
	jwtSecret        string
	accessExpirySec  int
	refreshExpirySec int
}

func NewAuthService(pool *pgxpool.Pool, jwtSecret string, accessExpiry, refreshExpiry int) *AuthService {
	return &AuthService{
		pool:             pool,
		jwtSecret:        jwtSecret,
		accessExpirySec:  accessExpiry,
		refreshExpirySec: refreshExpiry,
	}
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	ExpiresIn    int         `json:"expires_in"`
	User         *model.User `json:"user"`
}

func (s *AuthService) Register(ctx context.Context, input RegisterInput) (*model.User, error) {
	// Check if email exists
	var exists bool
	err := s.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 AND deleted_at IS NULL)", input.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailExists
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	user := &model.User{}
	err = s.pool.QueryRow(ctx,
		`INSERT INTO users (email, password_hash, name) VALUES ($1, $2, $3)
		 RETURNING id, email, name, created_at, updated_at`,
		input.Email, string(hash), input.Name,
	).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	user := &model.User{}
	err := s.pool.QueryRow(ctx,
		`SELECT id, email, password_hash, name, created_at, updated_at
		 FROM users WHERE email=$1 AND deleted_at IS NULL`,
		input.Email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return s.generateTokens(ctx, user)
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (*AuthResponse, error) {
	var userID string
	err := s.pool.QueryRow(ctx,
		`DELETE FROM refresh_tokens WHERE token=$1 AND expires_at > NOW() RETURNING user_id`,
		refreshToken,
	).Scan(&userID)
	if err != nil {
		return nil, ErrInvalidToken
	}

	user := &model.User{}
	err = s.pool.QueryRow(ctx,
		`SELECT id, email, password_hash, name, created_at, updated_at
		 FROM users WHERE id=$1 AND deleted_at IS NULL`,
		userID,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return s.generateTokens(ctx, user)
}

func (s *AuthService) ForgotPassword(ctx context.Context, email string) error {
	var userID string
	err := s.pool.QueryRow(ctx,
		`SELECT id FROM users WHERE email=$1 AND deleted_at IS NULL`,
		email,
	).Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}

	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return err
	}
	token := hex.EncodeToString(tokenBytes)

	expiresAt := time.Now().Add(1 * time.Hour)

	_, err = s.pool.Exec(ctx,
		`INSERT INTO password_reset_tokens (user_id, token, expires_at)
		 VALUES ($1, $2, $3)`,
		userID, token, expiresAt,
	)
	if err != nil {
		return err
	}

	log.Printf("Password reset token for %s: %s", email, token)

	return nil
}

func (s *AuthService) ResetPassword(ctx context.Context, token, newPassword string) error {
	var userID string
	err := s.pool.QueryRow(ctx,
		`UPDATE password_reset_tokens
		 SET used_at = NOW()
		 WHERE token=$1 AND expires_at > NOW() AND used_at IS NULL
		 RETURNING user_id`,
		token,
	).Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrTokenExpired
		}
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		return err
	}

	_, err = s.pool.Exec(ctx,
		`UPDATE users SET password_hash=$1, updated_at=NOW() WHERE id=$2`,
		string(hash), userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) generateTokens(ctx context.Context, user *model.User) (*AuthResponse, error) {
	now := time.Now()

	// Access token
	accessClaims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"iat":   now.Unix(),
		"exp":   now.Add(time.Duration(s.accessExpirySec) * time.Second).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSigned, err := accessToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	// Refresh token
	refreshBytes := make([]byte, 32)
	if _, err := rand.Read(refreshBytes); err != nil {
		return nil, err
	}
	refreshTokenStr := hex.EncodeToString(refreshBytes)

	_, err = s.pool.Exec(ctx,
		`INSERT INTO refresh_tokens (user_id, token, expires_at)
		 VALUES ($1, $2, $3)`,
		user.ID, refreshTokenStr, now.Add(time.Duration(s.refreshExpirySec)*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken:  accessSigned,
		RefreshToken: refreshTokenStr,
		ExpiresIn:    s.accessExpirySec,
		User:         user,
	}, nil
}
