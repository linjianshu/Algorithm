package main

import (
	"fmt"
	"sync"
	"time"
)

var A chan int
var B chan int
var wg sync.WaitGroup

func main() {
	A = make(chan int, 1) //这很重要!!!
	B = make(chan int, 1) //这很重要!!!

	wg.Add(1)
	A <- 1
	go func() {
		for {
			select {
			case <-A:
				fmt.Println("this is A")
				B <- 1
			default:
			}
		}
	}()
	time.Sleep(time.Second)
	go func() {
		for {
			select {
			case <-B:
				fmt.Println("this is B")
				A <- 1
			default:
			}
		}
	}()
	var flag string
	fmt.Scan(&flag)
	wg.Done()
	wg.Wait()

}
