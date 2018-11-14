package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	x := make(map[bool]int, 2)
	y := make(map[bool]int, 2)
	m := make(map[[2]int]bool, len(input))
	m[[2]int{0, 0}] = true
	for i := 0; i < len(input); i++ {
		e := i%2 == 0
		switch input[i] {
		case '^':
			y[e]--
		case '>':
			x[e]++
		case 'v':
			y[e]++
		case '<':
			x[e]--
		}
		m[[2]int{x[e], y[e]}] = true
	}
	fmt.Println(len(m))
}
