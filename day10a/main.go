package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func lookAndSay(str []byte) []byte {
	ch := str[0]
	count := 0
	result := []byte{}
	for i := 0; i < len(str); i++ {
		if str[i] != ch {
			result = append(result, byte(48+count), ch)
			ch = str[i]
			count = 0
		}
		count++
	}
	result = append(result, byte(48+count), ch)
	return result
}

func main() {
	input, _ := ioutil.ReadFile("input")
	input = []byte(strings.TrimSpace(string(input)))
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
		fmt.Printf("%v: %v\n", i, len(input))
	}
	fmt.Println(len(input))
}
