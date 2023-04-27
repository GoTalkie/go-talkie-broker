package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AuthServiceAddress = "auth-service"
	LoginEndpoint      = "/api/login"
	RegisterEndpoint   = "/api/register"
	ValidateEndpoint   = "/api/admin/validate"
)

func (app *Config) Auth(c *gin.Context) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s%s", AuthServiceAddress, ValidateEndpoint), c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer c.Request.Body.Close()
	req.Header = c.Request.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer resp.Body.Close()
	msg, _ := io.ReadAll(resp.Body)

	c.JSON(resp.StatusCode, gin.H{"message": msg})
}

func (app *Config) Login(c *gin.Context) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s%s", AuthServiceAddress, LoginEndpoint), c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer c.Request.Body.Close()
	req.Header = c.Request.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer resp.Body.Close()
	msg, _ := io.ReadAll(resp.Body)

	c.JSON(resp.StatusCode, gin.H{"message": msg})
}

func (app *Config) Register(c *gin.Context) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s%s", AuthServiceAddress, RegisterEndpoint), c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer c.Request.Body.Close()
	req.Header = c.Request.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer resp.Body.Close()
	msg, _ := io.ReadAll(resp.Body)

	c.JSON(resp.StatusCode, gin.H{"message": msg})
}
