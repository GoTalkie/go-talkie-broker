package server

import (
	"net/http"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

func (app *Config) Routes() http.Handler {
	r := gin.Default()

	r.POST("/handle", app.Handle)
	r.GET("/ping", healthcheck.Default())
	r.POST("/login", app.Login)
	r.POST("/register", app.Register)
	r.POST("/chat", app.Auth)

	return r
}
