package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		line := s.Text()
		cs := strings.Split(line, "x")
		c := make([]int, len(cs))
		for j := 0; j < len(cs); j++ {
			c[j], _ = strconv.Atoi(cs[j])
		}
		sort.Sort(sort.IntSlice(c))
		sum += 2*c[0] + 2*c[1] + c[0]*c[1]*c[2]

	}
	fmt.Println(sum)
}
