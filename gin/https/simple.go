/*
$GOROOT = go env GOROOT

windows:
go.exe run  $GOROOT\src\crypto\tls\generate_cert.go --host="localhost"

linux:
go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
*/

package main

import (
	"fmt"

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
	fmt.Printf("fuck")
}

func main() {
	server := gin.Default()
	server.GET("/", test)
	server.RunTLS("127.0.0.1:8888", "./cert.pem", "./key.pem")
}
