package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func func_get(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Method GET")
}

func func_post(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Method POST")
}

func main() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/GET", func_get).Methods("GET")
    myRouter.HandleFunc("/POST", func_post).Methods("POST")
    myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))


    fmt.Println("Listening on port 8888")
    log.Fatal(http.ListenAndServe("0.0.0.0:8888", myRouter))
}
