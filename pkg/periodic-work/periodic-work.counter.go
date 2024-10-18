package periodicwork

import (
	"fmt"
	"log"
	"sync"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type Counter struct {
	Started  int64
	Executed int64
	mu       sync.Mutex
}

func (c *Counter) IncStarted() {
	c.mu.Lock()
	c.Started++
	c.mu.Unlock()
}

func (c *Counter) IncExecuted() {
	c.mu.Lock()
	c.Executed++
	c.mu.Unlock()
}

func (c *Counter) Report() {
	c.mu.Lock()
	log.Printf("Counter Started: %d, Executed: %d\n", c.Started, c.Executed)
	c.mu.Unlock()
}

func (c *Counter) GetStatus() error {
	var status error = nil

	c.mu.Lock()
	if c.Started != c.Executed {
		status = ge.Pin(&ge.MismatchError{
			ComparedItems: fmt.Sprintf("Started: %d != Executed: %d\n", c.Started, c.Executed),
			Expected:      "Started == Executed",
			Actual:        "Started != Executed",
		})
	}
	c.mu.Unlock()

	return status
}
