package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AuthServiceAddress = "auth-service"
	LoginEndpoint      = "/api/login"
	RegisterEndpoint   = "/api/register"
	ValidateEndpoint   = "/api/admin/validate"
)

func (app *Config) Handle(c *gin.Context) {

}

func (app *Config) Auth(c *gin.Context) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s%s", AuthServiceAddress, ValidateEndpoint), c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	req.Header = c.Request.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp.Body)
}

func (app *Config) Login(c *gin.Context) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s%s", AuthServiceAddress, LoginEndpoint), c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	req.Header = c.Request.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp.Body)
}

func (app *Config) Register(c *gin.Context) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s%s", AuthServiceAddress, RegisterEndpoint), c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	req.Header = c.Request.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp.Body)
}
