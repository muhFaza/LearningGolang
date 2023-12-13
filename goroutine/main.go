package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int, 2)
	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		calculateSumOfSquares(numbers[5:], c)
		wg.Done()
	}()

	go func() {
		calculateSumOfSquares(numbers[:5], c)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	var temp int
	for v := range c {
		fmt.Println(v)
		temp += v
	}

	// time.Sleep(time.Second *2)
	fmt.Println(temp)
}

func calculateSumOfSquares (numbers []int, c chan int) {
	var temp int
	for _, v := range numbers{
		temp += (v * v)
	}
	c <- temp
}
