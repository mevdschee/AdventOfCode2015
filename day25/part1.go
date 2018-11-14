package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func coordToNumber(x, y int) int {
	n := 1
	for i := 0; i < x; i++ {
		n += i + 1
	}
	for i := 0; i < y+1; i++ {
		n += x + i
	}
	return n
}

func numberToCode(num int) int {
	n := 20151125
	for i := 0; i < num-1; i++ {
		n = (n * 252533) % 33554393
	}
	return n
}

func main() {
	input, _ := ioutil.ReadFile("input")
	re := regexp.MustCompile("^To continue, please consult the code grid in the manual.  Enter the code at row ([0-9]+), column ([0-9]+).$")
	parts := re.FindStringSubmatch(strings.TrimSpace(string(input)))
	row, _ := strconv.Atoi(parts[1])
	col, _ := strconv.Atoi(parts[2])
	fmt.Println(numberToCode(coordToNumber(col-1, row-1)))
}
