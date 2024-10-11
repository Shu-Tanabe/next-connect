package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/Shu-Tanabe/next-connect/go/internal-api/server"
	"github.com/Shu-Tanabe/next-connect/go/internal-api/utils"
)

func main() {

	// graceful shutdown
	ctx, stop := utils.GracefulShutdown()
	defer stop()

	defer utils.Logger.Sync()

	srv := server.NewServer()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				stackTrace := getStackTrace()
				utils.Logger.DPanic(fmt.Sprintf(
						"panic occured, %v, stacktrace: %s",
						r,
						stackTrace,
					),
				)
			}
		}()
		utils.Logger.Info(fmt.Sprintf("server started at %s", srv.Addr))
		err := srv.ListenAndServe()
		if err != nil {
			utils.Logger.DPanic(fmt.Sprintf("caught signal, %v", err))
		}
	}()
	<- ctx.Done()
	srv.Shutdown(ctx)
}

func getStackTrace() string {
	var stackTrace strings.Builder
	pc := make([]uintptr, 10)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])

	for {
		frame, more := frames.Next()
		stackTrace.WriteString(fmt.Sprintf("%s\t%s:%d", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}
	return stackTrace.String()
}