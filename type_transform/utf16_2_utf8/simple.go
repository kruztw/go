// ref: https://gist.github.com/bradleypeabody/185b1d7ed6c0c2ab6cec

package main

import (
	"bytes"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	b := []byte{
		0xff, // BOM
		0xfe, // BOM
		'T',
		0x00,
		'E',
		0x00,
		'S',
		0x00,
		'T',
		0x00,
		0x6C,
		0x34,
		'\n',
		0x00,
	}

	s1, err := DecodeUTF16_method1(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(s1)

	s2, err := DecodeUTF16_method2(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(s2)
}

func DecodeUTF16_method1(b []byte) (string, error) {
	if len(b)%2 != 0 || len(b) <= 1 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	b = b[2:] // ignore BOM
	u16s := make([]uint16, 1)
	ret := &bytes.Buffer{}
	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}

func DecodeUTF16_method2(b []byte) (string, error) {
	u8datas, _, err := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM).NewDecoder(), b)
	if err != nil {
		return "", err
	}

	return string(u8datas), nil
}
