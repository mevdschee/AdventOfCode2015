package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	target, _ := strconv.Atoi(strings.TrimSpace(string(input)))
	presents := make([]int, target+1)
	for e := 1; e <= target; e++ {
		delivered := 0
		for i := e; i <= target && delivered < 50; i += e {
			presents[i] += e * 11
			delivered++
		}
	}
	for i := 1; i <= target; i++ {
		if presents[i] >= target {
			fmt.Println(i)
			break
		}
	}
}
