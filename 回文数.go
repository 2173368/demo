package main

import "fmt"

func dem() {
	fmt.Println(isPalindrome(1598951))
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := 0
	y := x
	for {
		if y <= 0 {
			break
		}
		temp = y%10 + temp*10
		y = y / 10

	}
	return temp == x
}
