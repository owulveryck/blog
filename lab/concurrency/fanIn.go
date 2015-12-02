package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 11
	cs := make([]<-chan string, n)
	for i := 0; i < n; i++ {
		cs[i] = boring(fmt.Sprintf("Boring:%v", i))
	}
	c := fanIn(cs...)
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("This is the end!")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(1e3)) * 3 * time.Millisecond)
		c <- fmt.Sprintf("%v", msg)
	}()
	return c
}

func fanIn(input ...<-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			for _, n := range input {
				select {
				case s := <-n:
					c <- s
				default:
				}
			}
		}
	}()
	return c
}
