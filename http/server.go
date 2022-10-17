package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("received request: %v\n", r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`my first website`))
}

func main() {
	http.HandleFunc("/", test)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("ListenAndServe: %v\n", err)
	}
}
