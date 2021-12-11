package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func NewLifecycle() Lifecycle {
	return Lifecycle{}
}

type Lifecycle struct{}

func (l Lifecycle) WaitExitSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sigChan
	signal.Stop(sigChan)
}

func (l Lifecycle) Crash(err error) {
	logrus.WithError(err).Fatal("Lifecyclelication has crashed")
}
