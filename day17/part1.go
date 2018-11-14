package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func partUnit(prefix []int, number int, units []int, results *[][]int) {
	if number < 0 {
		return
	}
	if number == 0 {
		*results = append(*results, [][]int{prefix}...)
		return
	}
	n := len(units)
	for i := 0; i < n; i++ {
		u := units[i]
		newPrefix := append(append([]int{}, prefix...), u)
		newUnits := append([]int{}, units[i+1:n]...)
		partUnit(newPrefix, number-u, newUnits, results)
	}
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	units := []int{}
	for i := 0; s.Scan(); i++ {
		line := s.Text()
		u, _ := strconv.Atoi(line)
		units = append(units, u)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(units)))
	partitions := [][]int{}
	partUnit([]int{}, 150, units, &partitions)
	fmt.Println(len(partitions))
}
