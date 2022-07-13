/*Depth First Search*/

package main

import (
	"fmt"
)

type Node struct {
	Value	int
	Children 	[]*Node
}

func (n *Node) DFS(array []int) []int {
	array = append(array, n.Value)
	for _, child := range n.Children {
		array = child.DFS(array)
	}
}


