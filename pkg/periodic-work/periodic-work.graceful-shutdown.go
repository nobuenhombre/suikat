package periodicwork

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// kill (no param) default send syscall.SIGTERM
// kill -2 is syscall.SIGINT
// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
func (c *Config) waitOSInterruptSignal() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Periodic Work ...")
}

func (c *Config) GracefulShutDown() {
	c.waitOSInterruptSignal()
	c.Stop()
}
