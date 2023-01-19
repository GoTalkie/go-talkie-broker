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

	return r
}
