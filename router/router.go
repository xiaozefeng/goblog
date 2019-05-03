package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xiaozefeng/goblog/handler/blog"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(mw...)

	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OK\n")
	})

	g.GET("/", func(c *gin.Context) {
		articles, err := blog.ListBlog("assets/static/")
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    viper.GetString("blog_title"),
			"articles": articles,
		})
	})
	return g
}
