package main

import (
	"goin"
	"net/http"
)

func main() {
	g := goin.New()
	g.GET("/", func(ctx *goin.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Goin<h1>")
	})
	g.GET("/hello", func(ctx *goin.Context) {
		ctx.String(http.StatusOK, "hello %s\n", ctx.Query("name"))
	})
	g.POST("/login", func(ctx *goin.Context) {
		ctx.JSON(http.StatusOK, goin.H{
			"username": ctx.PostFrom("username"),
			"password": ctx.PostFrom("password"),
		})
	})
	g.Run(":9999")
}
