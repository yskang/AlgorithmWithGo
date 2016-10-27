package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var n, e, s int

	n = (rand.Int() % 500) + 1
	e = (rand.Int() % 1000) + 1
	s = (rand.Int() % 100) + 1

	fmt.Println("1")
	fmt.Printf("%d %d %d\n", n, e, s)
	for i := 1 ; i < n-1 ; i++ {
		fmt.Printf("%d %d %d\n", i, i+1, int32(rand.Int())%100)
	}

	for i := 0 ; i <= e - n ; i++ {
		var start, offset int32
		for {
			start = rand.Int31()%int32(n)+1
			if int32(n) - start != 0 {
				break
			}
		}

		offset =  rand.Int31()%(int32(n)-start)+1

		fmt.Printf("%d %d %d\n", start, start+offset, int32(rand.Int())%100)
	}

}
