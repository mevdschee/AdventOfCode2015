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
	lights := [1000][1000]int{}
	sum := 0
	for s.Scan() {
		line := s.Text()
		re1 := regexp.MustCompile("^(toggle|turn on|turn off)")
		operation := re1.FindString(line)
		re2 := regexp.MustCompile("([0-9]+)")
		cs := re2.FindAllString(line, 4)
		c := make([]int, len(cs))
		for i := 0; i < 4; i++ {
			c[i], _ = strconv.Atoi(cs[i])
		}
		for x := c[0]; x <= c[2]; x++ {
			for y := c[1]; y <= c[3]; y++ {
				switch operation {
				case "toggle":
					lights[x][y] += 2
				case "turn on":
					lights[x][y]++
				case "turn off":
					if lights[x][y] > 0 {
						lights[x][y]--
					}
				}
			}
		}
	}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			sum += lights[x][y]
		}
	}
	fmt.Println(sum)
}
