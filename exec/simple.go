package main

import (
	"fmt"
	"os/exec"
)

func main() {
    cmd := exec.Command("python3", "-V")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return
    }
    fmt.Println(string(out))
}
