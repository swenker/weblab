package main

import (
	"yinyushijing.cn/xnote-api/controller"
)

func main() {
	router := controller.SetupRouter()
	router.Run("localhost:8000")
}
