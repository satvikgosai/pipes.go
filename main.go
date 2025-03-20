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

const delay = time.Millisecond * 50

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

func getMultiplier() int {
	return rand.IntN(maxSize-minSize+1) + minSize
}

func main() {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Println("Error: Not running in a terminal")
		return
	}

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}
	height--

	matrix := make([][][]byte, height)
	for i := 0; i < height; i++ {
		matrix[i] = make([][]byte, width)
		for j := 0; j < width; j++ {
			matrix[i][j] = []byte(" ")
		}
	}

	var x, y int

	current := [2]int{0, 1}
	matrix[0][0] = horizontal

	for {
		x, y = (x+current[0]+height)%height, (y+current[1]+width)%width
		change := getDirection(current)
		matrix[x][y] = getAngle(current, change)
		current = change

		output := make([]byte, 0, len(matrix[0][0]))
		for _, row := range matrix {
			for _, chunk := range row {
				output = append(output, chunk...)
			}
		}
		os.Stdout.Write(clear)
		os.Stdout.Write(output)
		time.Sleep(delay)

		for i := 0; i <= getMultiplier(); i++ {
			x, y = (x+current[0]+height)%height, (y+current[1]+width)%width
			matrix[x][y] = straight[current]

			output := make([]byte, 0, len(matrix[0][0]))
			for _, row := range matrix {
				for _, chunk := range row {
					output = append(output, chunk...)
				}
			}
			os.Stdout.Write(clear)
			os.Stdout.Write(output)
			time.Sleep(delay)
		}
	}
}
