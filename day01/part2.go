package main

import "io/ioutil"
import "fmt"

func main() {
	input, _ := ioutil.ReadFile("input")
	sum := 0
	for i := 0; i < len(input); i++ {
		if sum < 0 {
			fmt.Println(i)
			break
		}
		switch input[i] {
		case '(':
			sum++
		case ')':
			sum--
		}
	}
}
