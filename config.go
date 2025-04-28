package main

import (
	"time"
)

const (
	MinSize  = 3
	MaxSize  = 5
	MinSpeed = 0
	MaxSpeed = 100
)

type PipeStyle struct {
	Horizontal string
	Vertical   string
	UpperRight string
	UpperLeft  string
	LowerRight string
	LowerLeft  string
}

const (
	Default = "\x1b[0m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Blue    = "\x1b[34m"
	Cyan    = "\x1b[36m"
	Magenta = "\x1b[35m"
	Yellow  = "\x1b[33m"
	Rainbow = "rainbow"
)

var (
	DefaultPipeStyle = PipeStyle{
		Horizontal: "═",
		Vertical:   "║",
		UpperRight: "╗",
		UpperLeft:  "╔",
		LowerRight: "╚",
		LowerLeft:  "╝",
	}
	SinglePipeStyle = PipeStyle{
		Horizontal: "─",
		Vertical:   "│",
		UpperRight: "┐",
		UpperLeft:  "┌",
		LowerRight: "└",
		LowerLeft:  "┘",
	}
	ThickPipeStyle = PipeStyle{
		Horizontal: "━",
		Vertical:   "┃",
		UpperRight: "┓",
		UpperLeft:  "┏",
		LowerRight: "┗",
		LowerLeft:  "┛",
	}
	RoundedPipeStyle = PipeStyle{
		Horizontal: "─",
		Vertical:   "│",
		UpperRight: "╮",
		UpperLeft:  "╭",
		LowerRight: "╰",
		LowerLeft:  "╯",
	}
	DottedPipeStyle = PipeStyle{
		Horizontal: "╾",
		Vertical:   "╿",
		UpperRight: "┑",
		UpperLeft:  "┍",
		LowerRight: "┕",
		LowerLeft:  "┙",
	}
)

var ColorThemes = map[string]string{
	"default": Default,
	"red":     Red,
	"green":   Green,
	"blue":    Blue,
	"cyan":    Cyan,
	"magenta": Magenta,
	"yellow":  Yellow,
	"rainbow": Rainbow,
}

var PipeStyles = map[string]PipeStyle{
	"default": DefaultPipeStyle,
	"single":  SinglePipeStyle,
	"thick":   ThickPipeStyle,
	"rounded": RoundedPipeStyle,
	"dotted":  DottedPipeStyle,
}

type Config struct {
	MinSize        int
	MaxSize        int
	UpdateInterval time.Duration
	ColorTheme     string
	PipeStyle      *PipeStyle
	RainbowMode    bool
}

func NewConfig(speed int, themeName string, styleName string) *Config {
	pipeStyle := PipeStyles[styleName]
	colorTheme := ColorThemes[themeName]
	rainbowMode := colorTheme == "rainbow"
	return &Config{
		MinSize:        MinSize,
		MaxSize:        MaxSize,
		UpdateInterval: time.Millisecond * time.Duration(100-speed),
		ColorTheme:     colorTheme,
		PipeStyle:      &pipeStyle,
		RainbowMode:    rainbowMode,
	}
}
