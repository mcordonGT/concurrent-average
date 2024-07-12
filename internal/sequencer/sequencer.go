package sequencer

import (
	"errors"
	"math/rand"
	"sync/atomic"
	"time"
)

type Sequencer interface {
	Next() (int, error)
	Average() float64
}

type sequencer struct {
	rand           *rand.Rand
	maxConcurrency int32
	sum            atomic.Int64
	count          atomic.Int32
	current        atomic.Int32
}

// New creates a new sequencer
func New() Sequencer {
	// Create a randomizer
	randomizer := rand.New(rand.NewSource(0))
	return &sequencer{
		rand:           randomizer,
		maxConcurrency: randomizer.Int31n(100) + 1,
		sum:            atomic.Int64{},
		count:          atomic.Int32{},
		current:        atomic.Int32{},
	}
}

// Next returns the next sequence number
func (s *sequencer) Next() (int, error) {
	// Limit the number of concurrent calls to Next
	if s.current.Load() > s.maxConcurrency {
		return 0, errors.New("cannot generate numbers concurrently")
	}

	// Mark that we are generating a number
	s.current.Add(1)
	defer s.current.Add(-1)

	// Wait for a random amount of time
	time.Sleep(time.Duration(s.rand.Intn(100)) * time.Millisecond)

	// Generate a random number
	n := s.rand.Intn(100)

	// Update the sum and count
	s.sum.Add(int64(n))
	s.count.Add(1)

	// Return the random number
	return n, nil
}

// Average returns the average of all the numbers generated
func (s *sequencer) Average() float64 {
	// Ensure that we are not generating numbers
	if s.current.Load() > 0 {
		panic("Cannot calculate average while generating numbers")
	}

	// Handle the case where no numbers have been generated
	if s.count.Load() == 0 {
		return 0
	}

	// Calculate the average
	return float64(s.sum.Load()) / float64(s.count.Load())
}
