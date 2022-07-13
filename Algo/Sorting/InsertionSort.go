/*Insertion Sort*/



package main

import "fmt"

func main() {
	var n = []int{43, 9, 7, 3, 4, 54, 11, 25}

	var i = 1
	for i < len(n) {
		var j = i 
		for j >= 1 && n[j] < n[j - 1] {
			n[j], n[j-1] = n[j-1], n[j]

			j--
		}
		i++
	}
	fmt.Println(n)
}