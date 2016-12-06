package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func perm(prefix []string, array []string, results *[][]string) {
	n := len(array)
	if n == 0 {
		*results = append(*results, [][]string{prefix}...)
	} else {
		for i := 0; i < n; i++ {
			newPrefix := append(append([]string{}, prefix...), array[i:i+1]...)
			newArray := append(append([]string{}, array[0:i]...), array[(i+1):n]...)
			perm(newPrefix, newArray, results)
		}
	}
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	scores := map[string]map[string]int{}
	for s.Scan() {
		line := s.Text()
		re := regexp.MustCompile("^([a-zA-Z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+).$")
		parts := re.FindStringSubmatch(line)
		a := parts[1]
		b := parts[4]
		d, _ := strconv.Atoi(parts[3])
		if parts[2] == "lose" {
			d *= -1
		}
		_, exists := scores[a]
		if !exists {
			scores[a] = make(map[string]int)
		}
		scores[a][b] = d
	}
	names := []string{}
	for name := range scores {
		names = append(names, name)
	}
	permutations := [][]string{}
	perm([]string{}, names, &permutations)
	first := true
	var highest int
	for i := 0; i < len(permutations); i++ {
		score := 0
		n := len(names)
		for p := 0; p < n; p++ {
			a := permutations[i][p]
			b := permutations[i][(p+1)%n]
			c := permutations[i][(p+n-1)%n]
			score += scores[a][b]
			score += scores[a][c]
		}
		if first || score > highest {
			highest = score
			first = false
		}
	}
	fmt.Println(highest)
}
