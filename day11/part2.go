package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func next(pwd []byte) []byte {
	for i := 7; i >= 0; i-- {
		pwd[i]++
		if pwd[i] > 122 {
			pwd[i] -= 26
		} else {
			break
		}
	}
	return pwd
}

func firstRule(pwd []byte) bool {
	for i := 0; i < len(pwd)-2; i++ {
		if pwd[i] == pwd[i+1]-1 && pwd[i] == pwd[i+2]-2 {
			return true
		}
	}
	return false
}

func secondRule(pwd []byte) bool {
	for i := 0; i < len(pwd); i++ {
		switch pwd[i] {
		case "i"[0]:
			return false
		case "o"[0]:
			return false
		case "l"[0]:
			return false
		}
	}
	return true
}

func thirdRule(pwd []byte) bool {
	pos := -1
	ch := byte(0)
	for i := 0; i < len(pwd)-1; i++ {
		if pos == -1 && pwd[i] == pwd[i+1] {
			pos = i
			ch = pwd[i]
		}
		if i > pos+1 && pwd[i] != ch && pwd[i] == pwd[i+1] {
			return true
		}
	}
	return false
}

func main() {
	input, _ := ioutil.ReadFile("input")
	pwd := []byte(strings.TrimSpace(string(input)))
	for {
		pwd = next(pwd)
		if firstRule(pwd) && secondRule(pwd) && thirdRule(pwd) {
			break
		}
	}
	for {
		pwd = next(pwd)
		if firstRule(pwd) && secondRule(pwd) && thirdRule(pwd) {
			break
		}
	}
	fmt.Println(string(pwd))
}
