package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalCatch(exitFunc func()) chan struct{} {
	sigChan := make(chan os.Signal)
	done := make(chan struct{})

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		exitFunc()
		close(done)
	}()

	return done
}
