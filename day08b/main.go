package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		line := s.Text()
		re1 := regexp.MustCompile("\\\\")
		new := re1.ReplaceAllLiteralString(line, "\\\\")
		re2 := regexp.MustCompile("\\\"")
		new = re2.ReplaceAllLiteralString(new, "\\\"")
		sum += 2 + len(new) - len(line)
	}
	fmt.Println(sum)
}
