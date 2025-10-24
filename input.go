package main

import (
	"os"
	"syscall"

	"golang.org/x/term"
)

var oldState *term.State

// SetupTerminal sets up the terminal for raw input
func SetupTerminal() error {
	var err error
	oldState, err = term.MakeRaw(int(syscall.Stdin))
	return err
}

// RestoreTerminal restores the terminal to normal mode
func RestoreTerminal() {
	if oldState != nil {
		term.Restore(int(syscall.Stdin), oldState)
	}
}

// GetKey gets a single key press
func GetKey() ([]byte, error) {
	b := make([]byte, 3)
	n, err := os.Stdin.Read(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
