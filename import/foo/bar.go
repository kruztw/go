package foo

import "fmt"

func World() string {               // 導出含式的首字母要大寫
    return " world"
}

func Hello() {
    fmt.Println("hello")
}
