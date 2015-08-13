package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"user"
	"webchat"
	"upload"
)

func main(){
	echo_serve := echo.New()

	mws := []echo.Middleware{mw.Logger(),mw.Recover()}
	//注册日志、故障恢复、响应Gzip压缩中间件
	echo_serve.Use(mws...)

	echo_serve.Static("/js","public/js")
	echo_serve.Static("/css","public/css")
	echo_serve.Static("/fonts","public/fonts")

	echo_serve.Static("/","templates")

	echo_serve.Post("/register",user.Register)
	echo_serve.Post("/upload",upload.UploadFile)

	echo_serve.WebSocket("/ws",webchat.WebChat)

	echo_serve.Run(":8080")
}