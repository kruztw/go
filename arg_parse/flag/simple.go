// flag 用來分析使用者輸入 (python 的 argparse)
// go run simple.go -h
// go run simple.go -u kruztw -h example

package main

import (
    "flag"
    "fmt"
)

func main()  {
    var user string
    var pwd string
    var host string
    var port int

    flag.StringVar(&user, "u", "", "用户名，默认为空")
    flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
    flag.StringVar(&host, "h", "localhost", "主机名，默认为 localhost")
    flag.IntVar(&port, "port", 3306, "duan端口号，默认3306")

    flag.Parse()
    fmt.Printf("\n user=%v \n pwd=%v \n host=%v \n port=%v \n", user, pwd, host, port)
}
