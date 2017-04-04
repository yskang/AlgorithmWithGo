package main

import (
	"fmt"
)

func main() {
	fmt.Println(countArrangement(6))
}

func countArrangement(N int) int {
	availables := make(map[int][]int)
	for i := 1  ; i <= N ; i++ {
		availables[i] = make([]int, 0)
		for j := 1 ; j <= N ; j++ {
			if j % i == 0 || i % j == 0 {
				availables[i] = append(availables[i], j)
			}
		}
	}

	tree := make([][]int, 0)
	for i := 0 ; i < len(availables[1]) ; i++ {
		tree = append(tree, []int{availables[1][i]})
	}

	for i := 2 ; i <= N ; i++ {
		treeSize := len(tree)
		for j := 0 ; j < treeSize ; j++ {
			for k := 0 ; k < len(availables[i]) ; k++ {
				for l := 0 ; l < len(tree[j]) ; l++ {
					if tree[j][l] == availables[i][k] {
						//l = len(tree[j])
						break
					}
					if l == len(tree[j]) - 1 {
						temp := make([]int, len(tree[j]))
						copy(temp, tree[j])
						temp = append(temp, availables[i][k])
						tree = append(tree, temp)
					}
				}
			}
		}

		tree = tree[treeSize:]
	}

	return len(tree)
}

