package main

import (
	"fmt"

	"golang.org/x/crypto/md4"
)

func main() {
	hash := md4.New()
	hash.Write([]byte("hello world"))
	digest := hash.Sum(nil)
	fmt.Printf("digest: \n%v\n", digest)
}
