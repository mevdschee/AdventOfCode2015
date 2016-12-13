package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(input), "\n")
	registers := map[string]int{"a": 1, "b": 0}
	pc := 0
	for pc < len(lines) {
		line := lines[pc]
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "hlf":
			registers[parts[1]] /= 2
		case "tpl":
			registers[parts[1]] *= 3
		case "inc":
			registers[parts[1]]++
		case "jmp":
			offset, _ := strconv.Atoi(parts[1])
			pc += offset - 1
		case "jie":
			register := strings.Trim(parts[1], ",")
			if registers[register]%2 == 0 {
				offset, _ := strconv.Atoi(parts[2])
				pc += offset - 1
			}
		case "jio":
			register := strings.Trim(parts[1], ",")
			if registers[register] == 1 {
				offset, _ := strconv.Atoi(parts[2])
				pc += offset - 1
			}
		}
		pc++
	}
	fmt.Println(registers["b"])
}
