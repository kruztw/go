package main

import (
    "fmt"
    "net/http"
    "log"
    "html/template"
    "encoding/json"
    "strings"
)

var tmpl = template.Must(template.ParseFiles("./login.html"))

func home(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query() // GET variable
    arg1, ok := query["arg1"]
    if ok && len(arg1[0]) > 0 {
        fmt.Printf("arg1 = %s\n", arg1[0])
    }
    fmt.Fprintf(w, "This is home")
}

func postVar(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    arg1:= strings.Join(r.PostForm["arg1"], "")
    if len(arg1) > 0 {
        fmt.Printf("arg1 = %s\n", arg1)
    }

    fmt.Fprintf(w, "This is postvar")
}

func home2(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Referer", "localhost")
    w.Write([]byte(`This is home2`))
}

func JSONmessage(w http.ResponseWriter, code int, msg string){
    json.NewEncoder(w).Encode(map[string]string{"msg": msg})
}

func home3(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"msg": "This is home3"})
    return
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
        obj := new(myTempl)
        obj.Account_str = "account"
        obj.Passwd_str  = "passwd"
        tmpl.Execute(w, obj)
    case "POST":
        cookie    :=    http.Cookie{Name:"user",Value:"guest"}
        http.SetCookie(w, &cookie)
        fmt.Fprintf(w, "login success")
    }
}


type XSS struct {
    Payload   template.HTML
}

func xssHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("./xss.html"))
    obj := XSS{Payload: template.HTML("<script>alert(1);</script>")}
    tmpl.Execute(w, obj)
}


func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/home2", home2)
    http.HandleFunc("/home3", home3)
    http.HandleFunc("/html", html)
    http.HandleFunc("/login", login)
    http.HandleFunc("/xss", xssHandler)
    http.HandleFunc("/postVar", postVar)

    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
