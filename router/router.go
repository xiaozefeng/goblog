package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(mw...)

	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OK\n")
	})
	g.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "index title",
		})
	})
	return g
}
