package main

import (
	"fmt"

	"github.com/mcordonGT/concurrent-average/internal/sequencer"
)

type Sequencer interface {
	Next() (int, error)
	Average() float64
}

func main() {
	// Create a new sequencer
	var s Sequencer = sequencer.New()

	// TODO: Replace dummy code below with your own code

	// Get the next number
	n, _ := s.Next()

	// Print the number and the average
	fmt.Println("n: ", n)

	// Get the average
	fmt.Println("s.Average(): ", s.Average())
}
