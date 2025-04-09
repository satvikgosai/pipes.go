package main

import (
	"os"
	"os/signal"
	"syscall"
)

func run(cfg *Config) error {
	term, err := NewTerminal()
	if err != nil {
		return err
	}

	term.Clear()
	term.HideCursor()
	defer term.ShowCursor()

	m := NewMatrix(term, cfg)

	// Handle window resize
	sigWinchChan := make(chan os.Signal, 1)
	signal.Notify(sigWinchChan, syscall.SIGWINCH)
	go func() {
		for range sigWinchChan {
			m.Reset()
		}
	}()

	// Handle termination
	sigTermChan := make(chan os.Signal, 1)
	signal.Notify(sigTermChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigTermChan:
			return nil
		default:
			m.Animate()
		}
	}
}

func main() {
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}
