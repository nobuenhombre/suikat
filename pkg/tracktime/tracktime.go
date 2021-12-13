// Package tracktime provides special struct to describe timer and log duration between start and stop points
//
// example:
//
// func SomeBodyDo() {
//     timer := tracktime.Start("Some Body Do")
//     defer func() {
//         timer.Stop()
//         timer.Log()
//     }()
//     ...
//     Here some actions
//     ...
// }
package tracktime

import (
	"log"
	"time"
)

// Tracker struct save info about timer
type Tracker struct {
	Label    string
	Run      time.Time
	Finish   time.Time
	Duration time.Duration
}

// Start create new Tracker
func Start(label string) *Tracker {
	return &Tracker{
		Label: label,
		Run:   time.Now(),
	}
}

// Stop fix time moment in Tracker
func (t *Tracker) Stop() {
	t.Finish = time.Now()
	t.Duration = time.Since(t.Run)
}

// Log logging Tracker info
func (t *Tracker) Log() {
	log.Printf(
		"track time:\n"+
			" - label [%v]\n"+
			" - run at [%v]\n"+
			" - finish at [%v]\n"+
			" - duration [%v]\n",
		t.Label,
		t.Run,
		t.Finish,
		t.Duration,
	)
}
