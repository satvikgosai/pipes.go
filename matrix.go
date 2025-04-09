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
		{1, 0}:  Vertical,
		{0, 1}:  Horizontal,
		{-1, 0}: Vertical,
		{0, -1}: Horizontal,
	}
	angles = map[[2]Direction]string{
		{{1, 0}, {0, 1}}:   LowerRight,
		{{1, 0}, {0, -1}}:  LowerLeft,
		{{-1, 0}, {0, 1}}:  UpperLeft,
		{{-1, 0}, {0, -1}}: UpperRight,
		{{0, 1}, {1, 0}}:   UpperRight,
		{{0, 1}, {-1, 0}}:  LowerLeft,
		{{0, -1}, {1, 0}}:  UpperLeft,
		{{0, -1}, {-1, 0}}: LowerRight,
	}
	redirect = map[Direction]Direction{
		{1, 0}:  {-1, 0},
		{0, 1}:  {0, -1},
		{-1, 0}: {1, 0},
		{0, -1}: {0, 1},
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
	matrix.data[0][0] = Horizontal
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
		m.data[m.x][m.y] = content
		m.terminal.WriteAt(m.x, m.y, content)
	}
	time.Sleep(cfg.UpdateInterval)
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
