// ref: https://ithelp.ithome.com.tw/articles/10205062

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:name`
}

func main() {
	data := []byte(`{"id" : 1 , "name" : "Daniel"}`)

	var person Person
	json.Unmarshal(data, &person)
	fmt.Println(person)

	jsondata, _ := json.Marshal(person)
	fmt.Println(string(jsondata))
}
