package main

import "time"

const (
	MinSize  = 3
	MaxSize  = 5
	MinSpeed = 0
	MaxSpeed = 100
)

const (
	Horizontal      = "═"
	Vertical        = "║"
	UpperRight      = "╗"
	UpperLeft       = "╔"
	LowerRight      = "╚"
	LowerLeft       = "╝"
	ClearScreen     = "\x1b[H\x1b[2J\x1b[3J"
	CursorInvisible = "\x1b[?25l"
	CursorVisible   = "\x1b[?12l\x1b[?25h"
	CursorMove      = "\x1b[%d;%dH%s"
)

type Config struct {
	MinSize        int
	MaxSize        int
	UpdateInterval time.Duration
	MaxSpeed       int
	MinSpeed       int
}

func NewConfig(speed int) *Config {
	return &Config{
		MinSize:        MinSize,
		MaxSize:        MaxSize,
		UpdateInterval: time.Millisecond * time.Duration(100-speed),
		MaxSpeed:       MaxSpeed,
		MinSpeed:       MinSpeed,
	}
}
