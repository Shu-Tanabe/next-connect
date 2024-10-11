package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown()  (ctxSig context.Context, stop context.CancelFunc) {
	ctxSig, stop = signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-ctxSig.Done()
		ctxTime, cancel := context.WithTimeout(context.Background(), time.Duration(30*time.Second))
		defer cancel()
		for {
			select {
				case <-ctxTime.Done():
					Logger.Error("force shutting down now...")
					cancel()
					os.Exit(1)
				case <- time.After(5*time.Second):
					Logger.Error("wait shutdown...")
			}
		}
	}()
	return ctxSig, stop
}