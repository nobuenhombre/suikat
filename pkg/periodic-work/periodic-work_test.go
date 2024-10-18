package periodicwork

import (
	"log"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func myRunTestWorker(wg *sync.WaitGroup, counter *Counter, config interface{}) {
	defer func() {
		counter.IncExecuted()
		wg.Done()
	}()

	time.Sleep(1 * time.Second)
	log.Printf("Executed some work  at %v\n", time.Now().UTC())
}

func TestRunWithGracefulShutDown(t *testing.T) {
	go func() {
		time.Sleep(300 * time.Millisecond)
		t.Logf("kill syscall.SIGTERM ...")
		err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		assert.NoError(t, err)
	}()

	config := &Config{
		WithGracefulShutDown: true,
		StopTimeout:          0 * time.Millisecond,
		WorkGapInterval:      100 * time.Millisecond,
		Worker:               myRunTestWorker,
	}

	pw := New(config)
	pw.Run()

	err := pw.GetStatus()
	assert.NoError(t, err)
}

func TestRunWithStop(t *testing.T) {
	config := &Config{
		WithGracefulShutDown: false,
		StopTimeout:          300 * time.Millisecond,
		WorkGapInterval:      100 * time.Millisecond,
		Worker:               myRunTestWorker,
	}

	pw := New(config)
	pw.Run()
	pw.Stop()

	err := pw.GetStatus()
	assert.NoError(t, err)
}
