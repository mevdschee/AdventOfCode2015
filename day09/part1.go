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
	distances := map[string]map[string]int{}
	for s.Scan() {
		line := s.Text()
		re := regexp.MustCompile("^([a-zA-Z]+) to ([a-zA-Z]+) = ([0-9]+)$")
		parts := re.FindStringSubmatch(line)
		a := parts[1]
		b := parts[2]
		d, _ := strconv.Atoi(parts[3])
		_, aExists := distances[a]
		if !aExists {
			distances[a] = make(map[string]int)
		}
		_, bExists := distances[b]
		if !bExists {
			distances[b] = make(map[string]int)
		}
		distances[a][b] = d
		distances[b][a] = d
	}
	cities := []string{}
	for city := range distances {
		cities = append(cities, city)
	}
	permutations := [][]string{}
	perm([]string{}, cities, &permutations)
	shortest := -1
	for i := 0; i < len(permutations); i++ {
		d := 0
		for p := 0; p < len(permutations[i])-1; p++ {
			a := permutations[i][p]
			b := permutations[i][p+1]
			d += distances[a][b]
		}
		if shortest == -1 || d < shortest {
			shortest = d
		}
	}
	fmt.Println(shortest)
}
