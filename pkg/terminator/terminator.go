package terminator

import (
	"github.com/nobuenhombre/suikat/pkg/terminator/terminated"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Terminator struct {
	Services []terminated.IGracefulShutDownService
}

type ITerminator interface {
	WaitOSInterruptSignalAndShutDown()
}

func NewTerminator(services []terminated.IGracefulShutDownService) ITerminator {
	return &Terminator{
		Services: services,
	}
}

// WaitOSInterruptSignalAndShutDown Wait for interrupt signal to gracefully shut down the services
// kill (no param) default send syscall.SIGTERM
// kill -2 is syscall.SIGINT
// kill -9 is syscall. SIGKILL but can"t be caught, so don't need to add it
func (t *Terminator) WaitOSInterruptSignalAndShutDown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown ...")

	for _, srv := range t.Services {
		srv.GracefulShutDown()
	}
}
