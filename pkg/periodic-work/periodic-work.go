package periodicwork

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type Worker func(wg *sync.WaitGroup, counter *Counter, config interface{})

type Config struct {
	WithGracefulShutDown bool
	StopTimeout          time.Duration
	WorkGapInterval      time.Duration
	Worker               Worker
	WorkerConfig         interface{}
	ctx                  context.Context
	executeWaitGroup     *sync.WaitGroup
	counter              *Counter
}

type Service interface {
	Run()
	Stop()
	GetStatus() error
}

func New(c *Config) Service {
	c.ctx = context.Background()

	return c
}

func (c *Config) execute() {
	defer c.executeWaitGroup.Done()

	wg := new(sync.WaitGroup)
	c.counter = &Counter{}

Exit:
	for {
		select {
		case <-c.ctx.Done():
			log.Printf("Stop at %v\n", time.Now().UTC())
			break Exit
		default:
			c.counter.IncStarted()
			log.Printf("Run at %v\n", time.Now().UTC())
			wg.Add(1)
			go c.Worker(wg, c.counter, c.WorkerConfig)
			time.Sleep(c.WorkGapInterval)
		}
	}

	wg.Wait()
	c.counter.Report()
}

func (c *Config) Run() {
	c.executeWaitGroup = new(sync.WaitGroup)

	go func() {
		c.executeWaitGroup.Add(1)
		c.execute()
	}()

	if c.WithGracefulShutDown {
		c.gracefulShutDown()
	}
}

func (c *Config) Stop() {
	ctx, cancel := context.WithTimeout(c.ctx, c.StopTimeout)
	defer cancel()

	c.ctx = ctx

	select {
	case <-c.ctx.Done():
		c.executeWaitGroup.Wait()
	}

	log.Println("Periodic Work exiting")
}

func (c *Config) GetStatus() error {
	err := c.counter.GetStatus()
	if err != nil {
		return ge.Pin(err)
	}

	return nil
}
