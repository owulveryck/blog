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
	n := 6
	cs := make([]<-chan Message, n)
	for i := 0; i < n; i++ {
		cs[i] = runme(fmt.Sprintf("[%v]", i), time.Duration(rand.Intn(1e3))*time.Millisecond)
	}

	for {
		for j := 0; j < n; j++ {
			select {
			case cmd := <-cs[j]:
				fmt.Println(cmd)
				// task j is done... Let's trigger the other one
			default:
			}
		}
	}
	fmt.Println("This is the end!")
}

func runme(cmd string, duration time.Duration) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool) // Shared between all messages.
	go func() {
		//for i := 0; ; i++ {
		//c <- Message{fmt.Sprintf("%s: %d", cmd, i), waitForIt}
		time.Sleep(duration)
		fmt.Printf("%v done\n", cmd)
		c <- Message{cmd, waitForIt}
		//<-waitForIt
		//}
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
