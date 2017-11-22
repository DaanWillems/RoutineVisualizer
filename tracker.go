package RoutineVisualizer

import (
	"encoding/json"
	"fmt"
	"time"
)

var trackers []Tracker

type Tracker struct {
	Name     string
	Timeline chan Event
	Start    time.Time
}

type Event struct {
	Name string
	Time float64
}

type TrackerResult struct {
	TrackerName string
	Events      []Event
}

func NewTracker(name string) Tracker {
	t := Tracker{
		Name:     name,
		Timeline: make(chan Event, 100),
		Start:    time.Now(),
	}
	trackers = append(trackers, t)
	return t
}

func WriteEvent(name string, t Tracker) {
	t.Timeline <- Event{
		Name: name,
		Time: time.Now().Sub(t.Start).Seconds(),
	}
}

func EncodeTrackers() {
	var results []TrackerResult
	for _, t := range trackers {
		close(t.Timeline)
		ev := TrackerResult{
			TrackerName: t.Name,
		}
		for event := range t.Timeline {
			ev.Events = append(ev.Events, event)
		}
		results = append(results, ev)
	}
	b, _ := json.Marshal(results)
	fmt.Println(string(b))
}
