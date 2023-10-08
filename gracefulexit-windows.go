//go:build windows
// +build windows

package gracefulexit

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signalHandler(signal os.Signal) {
	if signal == os.Interrupt || signal == syscall.SIGTERM {
		os.Exit(0)
	}
}

func HandleSignals(verbose bool) {
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, os.Interrupt)
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
