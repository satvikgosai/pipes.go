package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"golang.org/x/term"
)

const (
	minSize = 3
	maxSize = 5
)

var (
	horizontal = []byte("═")
	vertical   = []byte("║")
	upperRight = []byte("╗")
	upperLeft  = []byte("╔")
	lowerRight = []byte("╚")
	lowerLeft  = []byte("╝")
	clear      = []byte("\x1b[H\x1b[2J\x1b[3J")
)

var (
	directions = [4][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	straight = map[[2]int][]byte{
		{1, 0}:  vertical,
		{0, 1}:  horizontal,
		{-1, 0}: vertical,
		{0, -1}: horizontal,
	}
	angles = map[[2][2]int][]byte{
		{{1, 0}, {0, 1}}:   lowerRight,
		{{1, 0}, {0, -1}}:  lowerLeft,
		{{-1, 0}, {0, 1}}:  upperLeft,
		{{-1, 0}, {0, -1}}: upperRight,
		{{0, 1}, {1, 0}}:   upperRight,
		{{0, 1}, {-1, 0}}:  lowerLeft,
		{{0, -1}, {1, 0}}:  upperLeft,
		{{0, -1}, {-1, 0}}: lowerRight,
	}
	redirect = map[[2]int][2]int{
		{1, 0}:  {-1, 0},
		{0, 1}:  {0, -1},
		{-1, 0}: {1, 0},
		{0, -1}: {0, 1},
	}
)

func getMultiplier() int {
	return rand.IntN(maxSize-minSize+1) + minSize
}

func getDirection(current [2]int) [2]int {
	selected := directions[rand.IntN(4)]
	if current == redirect[selected] {
		return current
	}
	return selected
}

func getAngle(current [2]int, change [2]int) []byte {
	if current == change {
		return straight[current]
	}
	return angles[[2][2]int{current, change}]
}

type Matrix struct {
	x, y          int
	height, width int
	current       [2]int
	data          [][][]byte
	delay         time.Duration
}

func MatrixConstructor(height int, width int, speed int) *Matrix {
	data := make([][][]byte, height)
	for i := 0; i < height; i++ {
		data[i] = make([][]byte, width)
		for j := 0; j < width; j++ {
			data[i][j] = []byte(" ")
		}
	}
	data[0][0] = horizontal
	return &Matrix{
		height:  height,
		width:   width,
		current: [2]int{0, 1},
		data:    data,
		delay:   time.Millisecond * time.Duration(100-speed),
	}
}

func (m *Matrix) refresh() {
	matrix := m.data
	output := make([]byte, 0, len(matrix[0][0]))
	for _, row := range matrix {
		for _, chunk := range row {
			output = append(output, chunk...)
		}
	}
	os.Stdout.Write(clear)
	os.Stdout.Write(output)
	time.Sleep(m.delay)
}

func (m *Matrix) update(content []byte) {
	m.x, m.y = (m.x+m.current[0]+m.height)%m.height, (m.y+m.current[1]+m.width)%m.width
	m.data[m.x][m.y] = content
}

func run(speed int) error {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return fmt.Errorf("Not running in a terminal")
	}

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return fmt.Errorf("Unable to get terminal size: %v", err)
	}
	height--

	matrix := MatrixConstructor(height, width, speed)

	for {
		change := getDirection(matrix.current)
		matrix.update(getAngle(matrix.current, change))
		matrix.current = change
		matrix.refresh()
		for i := 0; i <= getMultiplier(); i++ {
			matrix.update(straight[matrix.current])
			matrix.refresh()
		}
	}
}
