package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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
	f, _ := os.Open("test")
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line + "=" + string(lookAndSay([]byte(line))))
	}

	input, _ := ioutil.ReadFile("input")
	input = []byte(strings.TrimSpace(string(input)))
	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
		fmt.Printf("%v: %v\n", i, len(input))
	}
	fmt.Println(len(input))
}
