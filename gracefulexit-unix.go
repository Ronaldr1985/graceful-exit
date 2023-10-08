//go:build linux || darwin
// +build linux darwin

package gracefulexit

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signalHandler(signal os.Signal) {
	if signal == syscall.SIGTERM || signal == syscall.SIGKILL || signal == syscall.SIGSTOP || signal == syscall.SIGINT {
		os.Exit(0)
	}
}

func HandleSignals(verbose bool) {
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, syscall.SIGINT)
	signal.Notify(signalChannel, syscall.SIGKILL)
	signal.Notify(signalChannel, syscall.SIGSTOP)
	signal.Notify(signalChannel, syscall.SIGTERM)

	go func() {
		var signal os.Signal
		for {
			signal = <-signalChannel

			if verbose {
				fmt.Println("signal:", signal)
			}

			signalHandler(signal)
		}
	}()
}
