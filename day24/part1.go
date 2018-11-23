package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func split(prefix1 []int, prefix2 []int, sum1, sum2, target int, numbers []int, results *[][2][]int) {
	if len(*results) > 0 {
		return
	}
	if sum1 > target || sum2 > target {
		return
	}
	if sum1 == target && sum2 == target {
		*results = append(*results, [2][]int{append([]int{}, prefix1...), append([]int{}, prefix2...)})
		return
	}
	if len(numbers) > 0 {
		number := numbers[0]
		newNumbers := append([]int{}, numbers[1:]...)
		split(append(prefix1, number), prefix2, sum1+number, sum2, target, newNumbers, results)
		split(prefix1, append(prefix2, number), sum1, sum2+number, target, newNumbers, results)
	}
}

func dfs(prefix []int, sum, best, target int, numbers []int, results *[][3][]int) {
	if sum > target {
		return
	}
	if sum == target {
		newNumbers := append([]int{}, numbers...)
		newPrefix := append([]int{}, prefix...)
		results2 := [][2][]int{}
		split([]int{}, []int{}, 0, 0, target, newNumbers, &results2)
		for _, result2 := range results2 {
			*results = append(*results, [3][]int{newPrefix, result2[0], result2[1]})
		}
		return
	}
	if len(prefix) < best {
		for i, number := range numbers {
			newNumbers := append(append([]int{}, numbers[:i]...), numbers[i+1:]...)
			newBest := len(prefix) + len(numbers)
			if len(*results) > 0 {
				newBest = len((*results)[0][0])
			}
			dfs(append(prefix, number), sum+number, newBest, target, newNumbers, results)
		}
	}
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	numbers := []int{}
	total := 0
	for s.Scan() {
		number, _ := strconv.Atoi(s.Text())
		numbers = append(numbers, number)
		total += number
	}
	part := total / 3
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	results := [][3][]int{}
	dfs([]int{}, 0, len(numbers), part, numbers, &results)
	lowest := -1
	for _, result := range results {
		product := 1
		for i := 0; i < len(result[0]); i++ {
			product *= result[0][i]
		}
		if lowest == -1 || product < lowest {
			lowest = product
		}
	}
	fmt.Println(lowest)
}
