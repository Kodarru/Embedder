package main

import (
	"fmt"
	"github.com/aymerick/raymond"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	HTML := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{title}}</title>
        <meta name="og:title" content="{{title}}">
		<meta name="og:description" content="{{description}}">
		<meta name="og:image" content="{{image}}">
		<meta name="og:url" content="{{url}}">
		<meta name="theme-color" content="{{color}}">

	</head>
	<body>
		<p>You weren't supposed to see this</p>
	</body>`
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		ctx := map[string]string{
			"title":       "",
			"description": "",
			"image":       "",
			"url":         "",
			"color":       "",
		}

		for key, value := range c.Request.URL.Query() {
			if len(value) > 0 {
				ctx[key] = value[0]
			}
		}

		result, err := raymond.Render(HTML, ctx)

		if err != nil {
			c.String(500, "Error rendering template: %s", err.Error())
			return
		}

		c.Header("Content-Type", "text/html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(result))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
