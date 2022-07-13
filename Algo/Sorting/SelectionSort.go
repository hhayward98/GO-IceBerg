/*SelectionSort*/
/*
Replace the current element in iteration by the lowest value on the right.
dont need to check last element because they way the algo sorts the left part of the array is already sorted.
*/


package main

import "fmt"

func main() {
	var n = []int{3, 52, 2, 54, 2, 34, 12, 9}

	var i = 1
	for i < len(n) - 1 {
		var j = i + 1
		var minIndex = i 

		if j < len(n) {
			if n[j] < n[minIndex] {
				minIndex = j
			}
			j++
		}

		if minIndex != i {
			var temp = n[i]
			n[i] = n[minIndex]
			n[minIndex] = temp
		}

		i++
	}

	fmt.Println(n)
}