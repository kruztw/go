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

func get(c *gin.Context) {
	var h header
	if err := c.ShouldBindHeader(&h); err != nil {
		c.String(404, "err: %v", err)
		return
	}

	fmt.Printf("h = %v\n", h)
	data := new(IndexData)
	data.Title = c.DefaultQuery("title", "default")
	data.Content = c.DefaultQuery("content", "default")
	c.HTML(http.StatusOK, "index.html", data)
}

func post(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		fmt.Printf("failed to get body: %v\n", err)
		c.Status(http.StatusNoContent)
		return
	}

	fmt.Printf("jsonData = %v\n", string(jsonData))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func get_post(c *gin.Context) {
	fmt.Printf("it works\n")
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.Use(gin.Recovery()) // (a)
	server.GET("/", get)
	server.POST("/post", post)
	server.GET("/panic", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil { // because of (a), return 200
				fmt.Printf("error: %v\n", err)
			}
		}()
		panic("panic")
	})

	server.Any("/any", get_post)
	server.Static("/wtf", "template/") // http://localhost:8888/wtf/index.html

	server.Run("127.0.0.1:8888")
}
