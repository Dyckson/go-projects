package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() (router *gin.Engine) {
	return setConfigs(gin.Default())
}

func setConfigs(router *gin.Engine) *gin.Engine {

	router.Use(cors.New(cors.Config{AllowOrigins: []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))

	router.Use(func(c *gin.Context) {
		header := c.Writer.Header()
		header.Add("Cache-Control", "no-cache, no-store, must-revalidate")
		header.Add("Pragma", "no-cache")
		header.Add("Expires", "0")
		c.Next()
	})

	return router
}
