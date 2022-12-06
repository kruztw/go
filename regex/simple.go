package main
 
import (
    "fmt"
    "regexp"
)
 
func main() {
    var word string

    fmt.Print("Enter any string: ")
    fmt.Scan(&word)

    is_alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(word)
    if is_alphanumeric{
        fmt.Printf("%s is an Alphanumeric string\n", word)
    } else{
        fmt.Printf("%s is not an Alphanumeric string\n", word)
    }
}
