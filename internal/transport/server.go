package transport

import (
	"net/http"

	_ "github.com/shop/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go Restful API with Swagger
// @version 1.0
// @description Simple swagger implementation in Go HTTP

// @contact.name Linggar Primahastoko
// @contact.url https://linggar.asia
// @contact.email x@linggar.asia

// @host localhost:8000
// @BasePath /
func ServerApp() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// http.HandleFunc("/auth/login", authLogin)
	// http.HandleFunc("/users/profile", userProfile)
	http.HandleFunc("/registration", Registration)

	http.ListenAndServe(":8000", nil)
}
