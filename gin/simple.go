package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

type header struct {
	Username string `header:"username"`
	Password string `header:"password`
}

func test(c *gin.Context) {
	var h header
	if err := c.ShouldBindHeader(&h); err != nil {
		c.String(404, "err: %v", err)
		return
	}

	fmt.Printf("h = %v\n", h)

	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.GET("/", test)
	server.Run("127.0.0.1:8888")
}
