package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			fmt.Println("hello world", i)
		}

	}()

	go func() {

		time.Sleep(1 * time.Millisecond)

		panic(errors.New("panic"))
		defer func() {
			i := recover()
			if i == errors.New("panic") {
				fmt.Println("go on")
			}
		}()
		defer wg.Done()

	}()
	wg.Wait()
	fmt.Println("over")
}
