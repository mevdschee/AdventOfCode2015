package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var wires = make(map[string]uint16)

func parseUint16(value string) (uint16, error) {
	number, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		wire, exists := wires[value]
		if !exists {
			return uint16(number), err
		}
		return wire, nil
	}
	return uint16(number), err
}

func run() {
	for {
		f, _ := os.Open("input")
		defer f.Close()
		s := bufio.NewScanner(f)
		for s.Scan() {
			line := s.Text()
			re1 := regexp.MustCompile("^(NOT )?([a-z0-9]+) -> ([a-z]+)$")
			parts1 := re1.FindStringSubmatch(line)
			if len(parts1) > 0 {
				if _, exists := wires[parts1[3]]; exists {
					continue
				}
				value, err := parseUint16(parts1[2])
				if err != nil {
					continue
				}
				if parts1[1] == "NOT " {
					value = ^value
				}
				wires[parts1[3]] = value
			}
			re2 := regexp.MustCompile("^([a-z0-9]+) (AND|OR|LSHIFT|RSHIFT) ([a-z0-9]+) -> ([a-z]+)$")
			parts2 := re2.FindStringSubmatch(line)
			if len(parts2) > 0 {
				if _, exists := wires[parts2[4]]; exists {
					continue
				}
				value1, err := parseUint16(parts2[1])
				if err != nil {
					continue
				}
				value2, err := parseUint16(parts2[3])
				if err != nil {
					continue
				}
				value := uint16(0)
				switch parts2[2] {
				case "AND":
					value = value1 & value2
				case "OR":
					value = value1 | value2
				case "LSHIFT":
					value = value1 << value2
				case "RSHIFT":
					value = value1 >> value2

				}
				wires[parts2[4]] = value
			}
		}
		_, done := wires["a"]
		if done {
			fmt.Println(wires["a"])
			break
		}
	}
}

func main() {
	run()
	a := wires["a"]
	wires = make(map[string]uint16)
	wires["b"] = a
	run()
}
