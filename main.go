package main

import (
	"net/http"

	"github.com/GrigorianNick/DnDCalendar/webserver"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func runWebserver() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	component := webserver.Index("Nick")
	router.GET("/test", gin.WrapH(templ.Handler(component)))
	router.Run()
}

func main() {
	runWebserver()
}
