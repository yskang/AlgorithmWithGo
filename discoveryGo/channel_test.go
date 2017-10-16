package discoveryGo

import (
	"fmt"
	"time"
	"sync"
)

func Example_simpleChannel() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// output:
	// 1
	// 2
	// 3
}

func Example_simpleChannel2() {
	c := make(chan int)
	go func() {
		c <- 4
		c <- 5
		c <- 6
		close(c)
	}()
	for num := range c {
		fmt.Println(num)
	}
	// output:
	// 4
	// 5
	// 6
}

func Example_simpleChannel3() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()

	for num := range c {
		fmt.Println(num)
	}

	// output:
	// 1
	// 2
	// 3
}

func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, ",")
	}
	// output:
	// 0,1,1,2,3,5,8,13,
}

func FibonacciGenerator(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}

func ExampleFibonacciGenerator() {
	fib := FibonacciGenerator(15)
	for n := fib() ; n >= 0; n = fib() {
		fmt.Print(n, ",")
	}
	// output:
	// 0,1,1,2,3,5,8,13,
}

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}

	// Output:
	// 7
	// 5
	// 10
}

type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func ExamplePlusTwo() {
	PlusTwo := Chain(PlusOne, PlusOne)
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 3
		c <- 6
		c <- 9
	}()
	for num := range PlusTwo(c) {
		fmt.Println(num)
	}
	// Output:
	// 5
	// 8
	// 11
}

func fanOut() {
	c := make(chan int)
	for i := 0 ; i < 3 ; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0 ; i < 10 ; i++ {
		c <- i
	}
	close(c)
}

func ExampleFanOut() {
	fanOut()
	// Output:
	//
}

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

