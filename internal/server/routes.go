package server

import (
	"net/http"

	"github.com/GoTalkie/go-talkie-broker/internal/middlewares"
	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

func (app *Config) Routes() http.Handler {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	r.GET("/ping", healthcheck.Default())
	r.POST("/login", app.Login)
	r.POST("/register", app.Register)

	protected := r.Group("/api")
	protected.Use(app.Auth)
	protected.POST("/chat", app.Chat)

	return r
}
