package RoutineVisualizer

import (
	"sync"
	"testing"
	"time"
)

func TestTracker(t *testing.T) {
	tracker1 := NewTracker("Worker 1")
	tracker2 := NewTracker("Worker 2")

	var wg sync.WaitGroup

	wg.Add(2)

	go func(t Tracker) {
		time.Sleep(1 * time.Second)
		WriteEvent("first event", t)
		time.Sleep(1 * time.Second)
		WriteEvent("second event", t)
		wg.Done()
	}(tracker1)

	go func(t Tracker) {
		time.Sleep(1 * time.Second)
		WriteEvent("first event", t)
		time.Sleep(1 * time.Second)
		WriteEvent("second event", t)
		wg.Done()
	}(tracker2)

	wg.Wait()

	EncodeTrackers()
}
