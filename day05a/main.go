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
		rule1, _ := regexp.MatchString("[aeiou].*[aeiou].*[aeiou]", line)
		//rule2, _ := regexp.MatchString("[a-z]\1", line)
		rule2 := false
		for i := 0; i < len(line)-1; i++ {
			if line[i] == line[i+1] {
				rule2 = true
			}
		}
		rule3, _ := regexp.MatchString("ab|cd|pq|xy", line)
		if rule1 && rule2 && !rule3 {
			sum++
		}
	}
	fmt.Println(sum)
}
