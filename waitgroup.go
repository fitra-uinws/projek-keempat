package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "manggo", "banana", "rambutan"}

	var wg sync.WaitGroup

	for i, fruit := range fruits {
		wg.Add(1)
		go printFruits(i, fruit, &wg)
	}

	wg.Wait()
}

func printFruits(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, fruit => %s \n", index, fruit)
	wg.Done()
}
