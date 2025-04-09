package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Terminal struct {
	Width  int
	Height int
}

func NewTerminal() (*Terminal, error) {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return nil, fmt.Errorf("not running in a terminal")
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, fmt.Errorf("unable to get terminal size: %w", err)
	}

	return &Terminal{
		Width:  width,
		Height: height,
	}, nil
}

func (t *Terminal) Clear() {
	fmt.Print("\x1b[H\x1b[2J\x1b[3J")
}

func (t *Terminal) HideCursor() {
	fmt.Print("\x1b[?25l")
}

func (t *Terminal) ShowCursor() {
	fmt.Print("\x1b[?12l\x1b[?25h")
}

func (t *Terminal) MoveCursor(x, y int) {
	fmt.Printf("\x1b[%d;%dH", x+1, y+1)
}

func (t *Terminal) WriteAt(x, y int, content string) {
	t.MoveCursor(x, y)
	fmt.Print(content)
}
