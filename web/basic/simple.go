package main

import (
    "fmt"
    "net/http"
    "log"
    "html/template"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is home")
}

func home2(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`This is home2`))
}

func html(w http.ResponseWriter, r *http.Request) {
    str := `<!DOCTYPE html>
<html>
<head></head>
<body><h1>html</h1></body>
</html>
`
    w.Write([]byte(str))
}

type myTempl struct {
    Account_str   string // 開頭要大寫
    Passwd_str    string
}

func login(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        tmpl := template.Must(template.ParseFiles("./login.html"))
        obj := new(myTempl)
        obj.Account_str = "account"
        obj.Passwd_str  = "passwd"
        tmpl.Execute(w, obj)
    case "POST":
        fmt.Fprintf(w, "login success")
    }
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/home2", home2)
    http.HandleFunc("/html", html)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
