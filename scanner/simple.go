package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

const delim = ","

func main() {
	input1 := "a b c"
	scanner1 := bufio.NewScanner(strings.NewReader(input1))

	scanner1.Split(bufio.ScanWords)
	for scanner1.Scan() {
		fmt.Println(scanner1.Text())
	}

	input2 := "1" + delim + "2" + delim + "3"
	scanner2 := bufio.NewScanner(strings.NewReader(input2))

	scanner2.Split(mySplit)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}
}

func mySplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte(delim)); i >= 0 {
		return i + len(delim), data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}
