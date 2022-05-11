package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// !Fix This Problem
func TestRaceCondition(t *testing.T) {
	var count int64

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				count++
			}
		}()
	}

	fmt.Println("Problem Race Condition: ", count)
}

// Solved Race Condition With Method One
func TestSolveRaceConditionMethodOne(t *testing.T) {
	var count int64
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Problem Race Condition: ", count)
}

// Solved Race Condition With Method Two
func TestSolveRaceConditionMethodTwo(t *testing.T) {
	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Problem Race Condition: ", count)
}
