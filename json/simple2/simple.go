// ref: https://ithelp.ithome.com.tw/articles/10205062

package main

import (
	"fmt"
)

type Friend struct {
	Name string
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:name`
	friends interface{} `json:friends`
}

func main() {
	var person Person = Person{
		Id: 1,
		Name: "xxx",
		friends: [...]Friend {
			{Name: "friend1"},
			{Name: "friend2"},
		},
	}

	fmt.Println(person)
}
