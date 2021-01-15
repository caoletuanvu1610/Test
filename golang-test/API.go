package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
)

func main() {

	ginS.GET("/product/:name", func(c *gin.Context) {

		c.String(200, c.Param("name"))
	})
	ginS.GET("/product/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := action + "-" + name
		c.String(200, message)
	})
	ginS.Run()
}
