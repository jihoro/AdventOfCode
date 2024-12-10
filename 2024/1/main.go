package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	var left, right []int
	count := make(map[int]int)
	for _, s := range strings.Split(string(data), "\n") {
		var n1, n2 int
		fmt.Sscanf(s, "%d   %d", &n1, &n2)
		left, right = append(left, n1), append(right, n2)
		count[n2]++
	}

	slices.Sort(left)
	slices.Sort(right)

	result, result2 := 0, 0
	for i := range left {
		result += abs(left[i] - right[i])
		result2 += left[i] * count[left[i]]
	}
	fmt.Println(result, result2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
