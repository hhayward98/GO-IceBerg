/*Merge Sort*/

/*
uses the divide and conquer method.
First divide the passed array by 2 recursively, until the length is 1 and then merge them.
*/

package main

import "fmt"

func merge(fp []int, sp []int) []int {
	var n = make([]int, len(fp)+len(sp))

	var fpIndex = 0
	var spIndex = 0
	var nIndex = 0

	for fpIndex < len(fp) && spIndex < len(sp) {
		if fp[fpIndex] < sp[spIndex] {
			n[nIndex] = fp[fpIndex]
			fpIndex++
		} else if sp[spIndex] < fp[fpIndex] {
			n[nIndex] = sp[spIndex]
			spIndex++
		}
		nIndex++
	}

	for fpIndex < len(fp) {
		n[nIndex] = fp[fpIndex]

		fpIndex++
		nIndex++
	}

	for spIndex < len(sp) {
		n[nIndex] = sp[spIndex]

		spIndex++
		nIndex++
	}

	return n
}

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	var fp = MergeSort(arr[0 : len(arr)/2])
	var sp = MergeSort(arr[len(arr)/2:])

	return merge(fp, sp)

}

func main() {
	 var n = []int{8, 4, 15, 7, 3, 13, 1, 5, 16}

	 fmt.Println(MergeSort(n))
}