package main

import "github.com/kkato/book-api/app/controller"

func main() {
	router := controller.GetRouter()
	router.Run() // 0.0.0.0:8080 でサーバーを立てる
}
