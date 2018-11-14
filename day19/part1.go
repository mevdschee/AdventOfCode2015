package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func replace(str, old, new string, next *map[string]bool) {
	p := len(str)
	for {
		p = strings.LastIndex(str[:p], old)
		if p < 0 {
			break
		}
		r := str[0:p] + new + str[p+len(old):]
		(*next)[r] = true
	}
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	trans := [][2]string{}
	molecules := map[string]bool{}
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
	molecule := s.Text()
	for t := range trans {
		replace(molecule, trans[t][0], trans[t][1], &molecules)
	}
	fmt.Println(len(molecules))
}
