package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func replace(str, old, new string, next *map[string]bool) bool {
	replaced := false
	p := len(str)
	for {
		p = strings.LastIndex(str[:p], old)
		if p < 0 {
			break
		}
		r := str[0:p] + new + str[p+len(old):]
		(*next)[r] = true
		replaced = true
	}
	return replaced
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	trans := [][2]string{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		re := regexp.MustCompile("^([a-zA-Z]+) => ([a-zA-Z]+)$")
		parts := re.FindStringSubmatch(line)
		trans = append(trans, [2]string{parts[1], parts[2]})
	}
	s.Scan()
	molecules := map[string]bool{s.Text(): true}
	target := "e"
	steps := 0
	for {
		steps++
		next := map[string]bool{}
		for molecule := range molecules {
			for t := range trans {
				if replace(molecule, trans[t][1], trans[t][0], &next) {
					break
				}
			}
		}
		molecules = next
		if molecules[target] {
			break
		}
	}
	fmt.Println(steps)
}
