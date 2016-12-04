package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		line := s.Text()
		rule1 := false
		for i := 0; i < len(line)-1; i++ {
			pair := string([]byte{line[i], line[i+1]})
			index := strings.Index(line[i+2:], pair)
			if index >= 0 {
				rule1 = true
				break
			}
		}
		rule2 := false
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				rule2 = true
				break
			}
		}
		if rule1 && rule2 {
			sum++
		}
	}
	fmt.Println(sum)
}
