package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    Id   int    `json:"id"`
    Name string `json:name`
}

func main() {
    data := []byte(`{"id" : 1 , "name" : "kruztw"}`)
    var person Person

    json.Unmarshal(data, &person) // string -> struct
    fmt.Println(person)

    jsondata, _ := json.Marshal(person) // struct -> string
    fmt.Println(string(jsondata))
}


