package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "time"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is home")
}


func main() {
    route := mux.NewRouter().StrictSlash(true)
    route.HandleFunc("/", home).Methods("GET")

    loggedRouter := handlers.LoggingHandler(os.Stdout, route)
    srv := &http.Server{
        Addr: "127.0.0.1:8888",
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler:      loggedRouter,
    }

    err := srv.ListenAndServe()
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
