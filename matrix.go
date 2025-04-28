package main

import (
	"math/rand/v2"
	"sync"
	"time"
)

type Direction [2]int

var (
	directions = [4]Direction{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	straight = map[Direction]string{
		{1, 0}:  DefaultPipeStyle.Vertical,
		{0, 1}:  DefaultPipeStyle.Horizontal,
		{-1, 0}: DefaultPipeStyle.Vertical,
		{0, -1}: DefaultPipeStyle.Horizontal,
	}
	angles = map[[2]Direction]string{
		{{1, 0}, {0, 1}}:   DefaultPipeStyle.LowerRight,
		{{1, 0}, {0, -1}}:  DefaultPipeStyle.LowerLeft,
		{{-1, 0}, {0, 1}}:  DefaultPipeStyle.UpperLeft,
		{{-1, 0}, {0, -1}}: DefaultPipeStyle.UpperRight,
		{{0, 1}, {1, 0}}:   DefaultPipeStyle.UpperRight,
		{{0, 1}, {-1, 0}}:  DefaultPipeStyle.LowerLeft,
		{{0, -1}, {1, 0}}:  DefaultPipeStyle.UpperLeft,
		{{0, -1}, {-1, 0}}: DefaultPipeStyle.LowerRight,
	}
	redirect = map[Direction]Direction{
		{1, 0}:  {-1, 0},
		{0, 1}:  {0, -1},
		{-1, 0}: {1, 0},
		{0, -1}: {0, 1},
	}
	rainbowColors = []string{
		Red,
		Yellow,
		Green,
		Cyan,
		Blue,
		Magenta,
	}
)

type Matrix struct {
	x, y          int
	height, width int
	current       Direction
	data          [][]string
	mutex         sync.Mutex
	terminal      *Terminal
	config        *Config
}

func NewMatrix(term *Terminal, cfg *Config) *Matrix {
	matrix := &Matrix{
		height:   term.Height,
		width:    term.Width,
		current:  Direction{0, 1},
		terminal: term,
		config:   cfg,
	}
	matrix.createEmptyData()
	matrix.data[0][0] = DefaultPipeStyle.Horizontal
	return matrix
}

func (m *Matrix) createEmptyData() {
	m.data = make([][]string, m.height)
	for i := 0; i < m.height; i++ {
		m.data[i] = make([]string, m.width)
	}
}

func (m *Matrix) Update(content string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.x, m.y = (m.x+m.current[0]+m.height)%m.height, (m.y+m.current[1]+m.width)%m.width

	if m.data[m.x][m.y] != content {
		colorCode := m.config.ColorTheme
		if m.config.RainbowMode {
			colorIndex := (m.x + m.y) % len(rainbowColors)
			colorCode = rainbowColors[colorIndex]
		}
		coloredContent := m.terminal.Colorize(content, colorCode)
		m.data[m.x][m.y] = content
		m.terminal.WriteAt(m.x, m.y, coloredContent)
	}

	time.Sleep(m.config.UpdateInterval)
}

func (m *Matrix) Reset() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if term, err := NewTerminal(); err == nil {
		m.terminal = term
		m.height = term.Height
		m.width = term.Width
		m.createEmptyData()
		m.terminal.Clear()
	}
}

func (m *Matrix) getMultiplier() int {
	return rand.IntN(m.config.MaxSize-m.config.MinSize+1) + m.config.MinSize
}

func (m *Matrix) getDirection() Direction {
	selected := directions[rand.IntN(4)]
	if m.current == redirect[selected] {
		return m.current
	}
	return selected
}

func (m *Matrix) getAngle(change Direction) string {
	if m.current == change {
		return straight[m.current]
	}
	return angles[[2]Direction{m.current, change}]
}

func (m *Matrix) Animate() {
	change := m.getDirection()
	m.Update(m.getAngle(change))
	m.current = change
	for i := 0; i <= m.getMultiplier(); i++ {
		m.Update(straight[m.current])
	}
}

func (m *Matrix) updatePipeCharacters() {
	style := m.config.PipeStyle
	straight = map[Direction]string{
		{1, 0}:  style.Vertical,
		{0, 1}:  style.Horizontal,
		{-1, 0}: style.Vertical,
		{0, -1}: style.Horizontal,
	}

	angles = map[[2]Direction]string{
		{{1, 0}, {0, 1}}:   style.LowerRight,
		{{1, 0}, {0, -1}}:  style.LowerLeft,
		{{-1, 0}, {0, 1}}:  style.UpperLeft,
		{{-1, 0}, {0, -1}}: style.UpperRight,
		{{0, 1}, {1, 0}}:   style.UpperRight,
		{{0, 1}, {-1, 0}}:  style.LowerLeft,
		{{0, -1}, {1, 0}}:  style.UpperLeft,
		{{0, -1}, {-1, 0}}: style.LowerRight,
	}
}
