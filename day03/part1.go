package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	x := 0
	y := 0
	m := make(map[[2]int]bool, len(input))
	m[[2]int{0, 0}] = true
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '^':
			y--
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		}
		m[[2]int{x, y}] = true
	}
	fmt.Println(len(m))
}
