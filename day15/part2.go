package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part(prefix []int, number int, parts int, results *[][]int) {
	if parts == 1 {
		newPrefix := append(append([]int{}, prefix...), number)
		*results = append(*results, [][]int{newPrefix}...)
	} else {
		for i := 0; i <= number; i++ {
			newPrefix := append(append([]int{}, prefix...), i)
			part(newPrefix, number-i, parts-1, results)
		}
	}
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	ingredients := map[int]ingredient{}
	for i := 0; s.Scan(); i++ {
		line := s.Text()
		re := regexp.MustCompile("^([a-zA-Z]+): capacity (-?[0-9]+), durability (-?[0-9]+), flavor (-?[0-9]+), texture (-?[0-9]+), calories (-?[0-9]+)$")
		parts := re.FindStringSubmatch(line)
		name := parts[1]
		capacity, _ := strconv.Atoi(parts[2])
		durability, _ := strconv.Atoi(parts[3])
		flavor, _ := strconv.Atoi(parts[4])
		texture, _ := strconv.Atoi(parts[5])
		calories, _ := strconv.Atoi(parts[6])
		ingredients[i] = ingredient{name, capacity, durability, flavor, texture, calories}
	}
	partitions := [][]int{}
	part([]int{}, 100, len(ingredients), &partitions)
	first := true
	var highest int
	for i := 0; i < len(partitions); i++ {
		capacity := 0
		durability := 0
		flavor := 0
		texture := 0
		calories := 0
		for p := 0; p < len(ingredients); p++ {
			capacity += partitions[i][p] * ingredients[p].capacity
			durability += partitions[i][p] * ingredients[p].durability
			flavor += partitions[i][p] * ingredients[p].flavor
			texture += partitions[i][p] * ingredients[p].texture
			calories += partitions[i][p] * ingredients[p].calories
		}
		score := 0
		if capacity > 0 && durability > 0 && flavor > 0 && texture > 0 && calories == 500 {
			score = capacity * durability * flavor * texture
		}
		if first || score > highest {
			highest = score
			first = false
		}
	}
	fmt.Println(highest)
}
