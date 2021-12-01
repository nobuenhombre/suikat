package tracktime

import (
	"log"
	"time"
)

//------------------------------------
// example:
// func SomeBodyDo() {
//     timer := tracktime.Start("Some Body Do")
//     defer func() {
//         timer.Stop()
//         timer.Log()
//     }()
//     ...
//
// }
//------------------------------------

type Tracker struct {
	Label    string
	Run      time.Time
	Finish   time.Time
	Duration time.Duration
}

func Start(label string) *Tracker {
	return &Tracker{
		Label: label,
		Run:   time.Now(),
	}
}

func (t *Tracker) Stop() {
	t.Finish = time.Now()
	t.Duration = time.Since(t.Run)
}

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
