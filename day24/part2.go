package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func split(prefix1 []int, prefix2 []int, prefix3 []int, sum1, sum2, sum3, target int, numbers []int, results *[][3][]int) {
	if len(*results) > 0 {
		return
	}
	if sum1 > target || sum2 > target || sum3 > target {
		return
	}
	if sum1 == target && sum2 == target && sum3 == target {
		*results = append(*results, [3][]int{append([]int{}, prefix1...), append([]int{}, prefix2...), append([]int{}, prefix3...)})
		return
	}
	if len(numbers) > 0 {
		number := numbers[0]
		newNumbers := append([]int{}, numbers[1:]...)
		split(append(prefix1, number), prefix2, prefix3, sum1+number, sum2, sum3, target, newNumbers, results)
		split(prefix1, append(prefix2, number), prefix3, sum1, sum2+number, sum3, target, newNumbers, results)
		split(prefix1, prefix2, append(prefix3, number), sum1, sum2, sum3+number, target, newNumbers, results)
	}
}

func dfs(prefix []int, sum, best, target int, numbers []int, results *[][4][]int) {
	if sum > target {
		return
	}
	if sum == target {
		newNumbers := append([]int{}, numbers...)
		newPrefix := append([]int{}, prefix...)
		results2 := [][3][]int{}
		split([]int{}, []int{}, []int{}, 0, 0, 0, target, newNumbers, &results2)
		for _, result2 := range results2 {
			*results = append(*results, [4][]int{newPrefix, result2[0], result2[1], result2[2]})
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
	part := total / 4
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	results := [][4][]int{}
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
