package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/xiaozefeng/goblog/router"
	"github.com/xiaozefeng/wbdc/config"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	cfg = flag.String("c", "", "blog server config")
)

func main() {
	err := config.Init(*cfg)
	if err != nil {
		panic(err)
	}

	// gin
	gin.SetMode(viper.GetString("runmode"))
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	g := gin.New()
	gin.DisableConsoleColor()
	g.Static("/assets", "./assets")
	g.StaticFile("/favicon.ico", "./resources/favicon.ico")
	g.LoadHTMLGlob("templates/*")
	var mw []gin.HandlerFunc

	router.Load(g, mw...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Printf("The router has been deployed successlly")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("port"))
	log.Println(http.ListenAndServe(viper.GetString("port"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		//ping the server by sending a get request to `/ping`
		resp, err := http.Get(viper.GetString("url") + "/ping")
		if err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Printf("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("can not connect the router")
}
