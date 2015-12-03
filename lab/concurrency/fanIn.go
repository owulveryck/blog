package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan mat64.Dense
}

func main() {
	// Allocate a zeroed array of size 8×8
	m := mat64.NewDense(8, 8, nil)
	m.Set(0, 1, 1)
	m.Set(0, 4, 1) // First row
	m.Set(1, 6, 1)
	m.Set(1, 6, 1) // second row
	m.Set(3, 2, 1)
	m.Set(3, 6, 1) // fourth row
	m.Set(5, 0, 1)
	m.Set(5, 1, 1)
	m.Set(5, 2, 1) // fifth row
	m.Set(7, 6, 1) // seventh row
	fa := mat64.Formatted(m, mat64.Prefix("    "))
	// Display the matrix
	fmt.Printf("\nm = %v\n\n", fa)

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
	waitForIt := make(chan mat64.Dense) // Shared between all messages.
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
