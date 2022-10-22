package main

import (
	"fmt"
	"time"
)

func SliceIterChan(s []int) <-chan int {
	outChan := make(chan int)
	go func() {
		for i := range s {
			outChan <- s[i]
		}
		close(outChan)
	}()
	return outChan
}

/*
	if you need to use range to pull channel data you need to know when the channel close
*/
func pannicFn1() {
	biDir := make(chan string)
	//everything is synchronized
	biDir <- "hello"
	close(biDir)
	// panic! because channel is closed nothing to read
	for i := range biDir {
		// fmt.Println(<-biDir)
		fmt.Println(i)
	}
}

func pannicFn2() {
	biDir := make(chan string)
	biDir <- "hello"
	// painc! because channle never closed
	for i := range biDir {
		// fmt.Println(<-biDir)
		fmt.Println(i)
	}
}

func pannicFn3() {
	biDir := make(chan string)
	go func() {
		biDir <- "hello"
	}()

	// painc! because channle never closed. the range biDir nerver be terminated
	for i := range biDir {
		// fmt.Println(<-biDir)
		fmt.Println(i)
	}
}

func pannicFn4() {
	// buffered with 1
	biDir := make(chan string, 1)
	go func() {
		biDir <- "hello"
	}()

	// works. the ^^ gorouting run background range biDir need be terminated
	for i := range biDir {
		// fmt.Println(<-biDir)
		fmt.Println(i)
	}
}

func pannicFn5() {
	biDir := make(chan string, 2)
	go func() {
		biDir <- "hello"
		//biDir <- "world"
		close(biDir)
	}()

	loop := make([]int, 2)
	for i := range loop {
		fmt.Println("oo", i)
		fmt.Println(<-biDir)
	}
}

func pannicFn6() {
	biDir := make(chan string)
	go func() {
		biDir <- "hello"
		close(biDir)
	}()

	// works. the ^^ gorouting run background range biDir need be terminated
	for i := range biDir {
		fmt.Println(i)
	}
}

func main() {
	//pannicFn3()
	// s := []int{1, 2, 3, 4, 5}
	// c := SliceIterChan(s)
	// for i := range c {
	// 	fmt.Println(i)
	// }

	biDir := make(chan string)
	// go func() {
	// 	time.Sleep(1*time.Second)
	// 	biDir <- "hello"
	// }()
	go pop(biDir)

	// painc! because channle never closed. the range biDir nerver be terminated
	for i := range biDir {
		// fmt.Println(<-biDir)
		fmt.Println(i)
	}
}

func pop(biDir chan string){
	time.Sleep(1*time.Second)
	biDir <- "hello"
}
