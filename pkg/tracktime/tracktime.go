package tracktime

import (
	"log"
	"time"
)

//------------------------------------
// example:
// func SomeBodyDo() {
//     defer tracktime.Stop(tracktime.Start("SomeBodyDo"))
//
//     ...
//
// }
//------------------------------------

func Start(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Stop(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
