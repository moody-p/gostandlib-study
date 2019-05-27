package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int{
	lookup := make(map[int]int)
	for i, v := range nums{
		j, ok := lookup[target-v]
		lookup[v] = i
		if ok {
			return []int{j, i}
		}
	}
	return []int{}
}
func main(){
	a := []int{2,3,5,9,19,25,67}
	target := 69
	fmt.Println(twoSum(a, target))
}