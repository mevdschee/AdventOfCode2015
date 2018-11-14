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
		re := regexp.MustCompile("(\\\\\\\\)|(\\\\\\\")|(\\\\x[0-9a-f]{2})")
		matches := re.FindAllStringSubmatch(line, -1)
		sum += 2
		count := len(matches)
		for i := 0; i < count; i++ {
			sum += len(matches[i][0]) - 1
		}
	}
	fmt.Println(sum)
}
