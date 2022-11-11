package main
import (
    "fmt"
    "os/exec"
    "bytes"
)

// /usr/bin/osascript -e 'tell app "System Events" to display dialog "Hello World"'

func main() {
    cmd := exec.Command("osascript", "-e", `tell app "System Events" to display dialog "Loggout"`)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        return
    }
    fmt.Println("Result: " + out.String())
}
