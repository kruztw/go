// ref: https://gist.github.com/bradleypeabody/185b1d7ed6c0c2ab6cec

package main

import "fmt"
import "unicode/utf16"
import "unicode/utf8"
import "bytes"

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

	s, err := DecodeUTF16(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}

func DecodeUTF16(b []byte) (string, error) {
	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

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
