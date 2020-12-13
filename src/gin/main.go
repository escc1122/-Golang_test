package main

import (
	"Golang_test/src/gin/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	apiServer := gin.New()
	apiServer.LoadHTMLGlob("src/gin/view/*")

	apiServer.Static("/assets", "./src/gin/view")

	apiServer.GET("/hello", HelloWorld)

	apiServer.Run(":8080")
}

func HelloWorld(context *gin.Context) {
	b, _ := json.Marshal(service.GetStockDay())
	c := string(b)
	fmt.Print(c)
	context.JSON(http.StatusOK, service.GetStockDay())
	//context.HTML(http.StatusOK, "index.html","")
}
