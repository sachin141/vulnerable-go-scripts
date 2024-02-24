package main

import (
	"fmt"
	"sync"
	"time"
)

type KeyStore struct {
	mu       sync.Mutex
	unlocked map[string]string
}

type common struct {
	Address string
}

func (ks *KeyStore) Lock(addr common) error {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	if unl, found := ks.unlocked[addr.Address]; found {
		ks.expire(addr.Address, unl, time.Duration(0)*time.Nanosecond)
	} else {
		ks.unlocked[addr.Address] = "NewValue" // Simulate modification of shared data
	}
	return nil
}

func (ks *KeyStore) expire(addr, unl string, duration time.Duration) {
	// Simulating expire function
	fmt.Printf("Expired: %v - %v - %v\n", addr, unl, duration)
}

func main() {
	var wg sync.WaitGroup
	ks := &KeyStore{
		unlocked: make(map[string]string),
	}

	// Simulating concurrent access to Lock function
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ks.Lock(common{Address: "example"})
		}(i)
	}

	wg.Wait()
	time.Sleep(1 * time.Second) // Keep program alive to observe the output
}
