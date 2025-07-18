package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range nums {
               if idx, ok := seen[target - v]; ok {
                   return []int{idx, i}
                }
                seen[v] = i
         }
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
