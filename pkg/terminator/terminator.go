package terminator

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nobuenhombre/suikat/pkg/terminator/terminated"
)

type Service struct {
	Services []terminated.IGracefulShutDownService
}

type ITerminator interface {
	WaitOSInterruptSignalAndShutDown()
}

func New(services []terminated.IGracefulShutDownService) ITerminator {
	return &Service{
		Services: services,
	}
}

// WaitOSInterruptSignalAndShutDown Wait for interrupt signal to gracefully shut down the services
// kill (no param) default send syscall.SIGTERM
// kill -2 is syscall.SIGINT
// kill -9 is syscall. SIGKILL but can"t be caught, so don't need to add it
func (t *Service) WaitOSInterruptSignalAndShutDown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown ...")

	for _, s := range t.Services {
		s.GracefulShutDown()
	}
}
