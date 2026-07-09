package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

// SwaggerHandler serves the swagger docs directory.
func SwaggerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		swaggerDir := os.Getenv("SWAGGER_DIR")
		if swaggerDir == "" {
			swaggerDir = "./docs"
		}
		http.StripPrefix("/swagger/", http.FileServer(http.Dir(filepath.Join(swaggerDir, "swagger")))).ServeHTTP(w, r)
	}
}
