/*Bubble Sort*/


package main

import "fmt"

func main() {
	var n = []int{2, 34, 12, 7, 4, 14, 9, 3}

	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n) - 1 {
			if n[i], n[i + 1]{
				n[i], n[i + 1] = n[i + 1], n[i]
				isDone = false
				
			}
			i++
		}
	}

	fmt.Println(n)
}



