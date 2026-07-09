package handler

import "net/http"

func SwaggerHandler() http.HandlerFunc {
	return http.FileServer(http.Dir("./docs/swagger/")).ServeHTTP
}
