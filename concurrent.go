package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("main execution started")
	go firstProcess(8)

	secondProcess(8)
	fmt.Println("No Of Goroutine ", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)
	fmt.Println("Main Execution ended")
}

func firstProcess(index int) {
	fmt.Println("First Process funct started")
	for i := 1; i <= index; i++ {
		fmt.Println("first i=", i)
	}
	fmt.Println("First Process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second Process funct started")
	for j := 1; j <= index; j++ {
		fmt.Println("second j=", j)
	}
	fmt.Println("Second Process func ended")
}
