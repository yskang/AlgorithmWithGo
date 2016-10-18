// https://algospot.com/judge/problem/read/HAMMINGCODE

package main

import (
	"fmt"
)

func main() {
	var numOfTest int
	fmt.Scanf("%d\n", &numOfTest)

	for i := 0 ; i < numOfTest ; i++ {
		var bit string
		bits := make([]int, 7)
		fmt.Scanf("%s\n", &bit)
		for j, b := range bit {
			bits[j] = int (b - '0')
		}

		p := make([]int, 3)

		p[0] = ( bits[2] + bits[4] + bits[6] + bits[0] ) % 2
		p[1] = ( bits[2] + bits[5] + bits[6] + bits[1] ) % 2
		p[2] = ( bits[4] + bits[5] + bits[6] + bits[3] ) % 2

		errorCode := p[2]<<2 + p[1]<<1 + p[0]

		if errorCode != 0 {
			bits[errorCode-1] ^= 1
		}

		out := string([]byte{byte('0' + bits[2]), byte('0' + bits[4]), byte('0' + bits[5]), byte('0' + bits[6])})
		fmt.Println(out)
	}
}
