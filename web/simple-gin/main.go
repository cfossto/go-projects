package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){

	// Default router.
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	
	// Get-endpoint.
	r.GET("/hello", func(c *gin.Context) {
		// Make endpoint return status 200 and return a Go map (key/value)
		c.JSON(200, gin.H{
			"msg":"working",
		})
	})

	r.GET("/site", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,"index.tmpl", gin.H{
			"firstHeading":"Welcome",
		},)
	})

	r.Run()
}
