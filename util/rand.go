// Package util provides public goroutine-safe random function.
// The implementation is similar to grpc random functions. Additionally,
// the seed function is provided to be called from the outside, and
// the random functions are provided as a body's methods.
package util

import (
	"math/rand"
	"sync"
)

type SafeRand struct {
	r  *rand.Rand
	mu sync.Mutex
}

func NewSafeRand(seed int64) *SafeRand {
	c := &SafeRand{
		r: rand.New(rand.NewSource(seed)),
	}
	return c
}

func (c *SafeRand) Intn(n int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	res := c.r.Intn(n)
	return res
}
