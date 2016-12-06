package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	aunts := []map[string]int{}
	for s.Scan() {
		line := s.Text()
		re := regexp.MustCompile("([a-z]+): ([0-9]+)")
		parts := re.FindAllStringSubmatch(line, -1)
		a := map[string]int{}
		for i := 0; i < len(parts); i++ {
			key := parts[i][1]
			val, _ := strconv.Atoi(parts[i][2])
			a[key] = val
		}
		aunts = append(aunts, a)
	}
	aunt := map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1}
	sue := 0
	for i := 0; i < len(aunts); i++ {
		match := true
		for prop := range aunt {
			val, exists := aunts[i][prop]
			if exists && val != aunt[prop] {
				match = false
			}
		}
		if match {
			sue = i + 1
			break
		}
	}
	fmt.Println(sue)
}
