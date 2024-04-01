package helpers

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal() os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	return <-stop
}
