package main

import "fmt"

func main2() {
	var nums = [7]int{1, 2, 3, 4, 4, 3, 1}
	fmt.Println(singleNumber(nums[:]))

}

func singleNumber(nums []int) int {
	record := make(map[int]int, 7)

	for i := range nums {
		element := nums[i]
		recordValue, exist := record[element]
		if exist {
			record[element] = recordValue + 1
		} else {
			record[element] = 1
		}
	}
	for key, value := range record {
		if value == 1 {
			return key
		}
	}
	return 0
}
