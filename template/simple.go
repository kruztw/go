package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"text/template"

	"golang.org/x/text/encoding/unicode"
)

type info struct {
	Name        string
	Description string
}

func main() {
	data := info{"user1", "the first user"}

	t, err := template.New("user_templ").Parse("\"{{ .Name}}\" description: \"{{ .Description}}\"")
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n ---------------------------- Encode result to base64 ---------------------------- \n")
	var out bytes.Buffer
	b64Encoder := base64.NewEncoder(base64.StdEncoding, &out)
	utf16Encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	utf16Writer := utf16Encoder.Writer(b64Encoder)

	if err := t.Execute(utf16Writer, data); err != nil {
		panic(err)
	}

	fmt.Printf("base64 result: %v\n", out.String())
}
