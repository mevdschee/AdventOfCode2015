package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	field := [100][100]bool{}
	for y := 0; s.Scan(); y++ {
		line := s.Text()
		for x := 0; x < 100; x++ {
			if line[x] == '#' {
				field[x][y] = true
			}
		}
	}
	for i := 0; i < 100; i++ {
		next := [100][100]bool{}
		next[0][0] = true
		next[99][0] = true
		next[0][99] = true
		next[99][99] = true
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				neighbors := 0
				for dx := x - 1; dx <= x+1; dx++ {
					for dy := y - 1; dy <= y+1; dy++ {
						if dx == x && dy == y {
							continue
						}
						if dx < 0 || dx >= 100 {
							continue
						}
						if dy < 0 || dy >= 100 {
							continue
						}
						if field[dx][dy] {
							neighbors++
						}
					}
				}
				if field[x][y] {
					if neighbors == 2 || neighbors == 3 {
						next[x][y] = true
					}
				} else {
					if neighbors == 3 {
						next[x][y] = true
					}
				}
			}
		}
		field = next
	}
	sum := 0
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			if field[x][y] {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
