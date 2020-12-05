package main

import (
	"sort"
)

func minMoves(nums []int) int {
	sort.Ints(nums)
	len := len(nums)
	m := nums[len/2]
	var sum int = 0
	for i:=0; i < len; i++ {
		sum = sum + abs((nums[i]-m))
	}
	return sum
}

func abs (i int) int {
	if i>0 {
		i = -i
	}
	return i
}

func main() {
	m :=[]int{1,2,3}
	minMoves(m)
}
