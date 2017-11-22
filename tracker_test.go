package RoutineVisualizer

import (
	"testing"
	"time"
)

func TestTracker(t *testing.T) {
	tracker1 := NewTracker("Worker 1")
	tracker2 := NewTracker("Worker 2")

	go func(t Tracker) {
		time.Sleep(1 * time.Second)
		WriteEvent("first", t)
		time.Sleep(1 * time.Second)
		WriteEvent("second", t)
	}(tracker1)

	go func(t Tracker) {
		time.Sleep(1 * time.Second)
		WriteEvent("first", t)
		time.Sleep(1 * time.Second)
		WriteEvent("second", t)
	}(tracker2)

	time.Sleep(4 * time.Second)

	EncodeTrackers()
}
