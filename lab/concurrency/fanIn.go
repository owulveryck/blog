package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	n := 3
	cs := make([]<-chan Message, n)
	for i := 0; i < n; i++ {
		cs[i] = boring(fmt.Sprintf("Boring:%v", i))
	}

	c := fanIn(cs...)
	for i := 0; i < n; i++ {
		msg := <-c
		msg.wait <- true
		fmt.Println(msg.str)

	}
	fmt.Println("This is the end!")
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool) // Shared between all messages.
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
			fmt.Printf("%v: %d WAITING\n", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			fmt.Printf("%v: %d DONE\n", msg, i)
			<-waitForIt
		}
	}()
	return c
}
func fanIn(inputs ...<-chan Message) <-chan Message { // HL
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}
