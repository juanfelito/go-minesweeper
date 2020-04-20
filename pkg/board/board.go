package board

import (
	"fmt"
	"go-minesweeper/pkg/cell"
	"math/rand"
	"time"
)

type Board struct {
	Cells  [][]*cell.Cell
	Values [][]int
	width  int
	height int
	mines  int
}

func init() {
	rand.Seed(time.Now().Unix())
}

func (b *Board) PrintBoard() {
	for _, i := range b.Values {
		for _, j := range i {
			fmt.Printf("%v ", j)
		}
		fmt.Print("\n")
	}
}

func (b *Board) placeMines() {
	for i := 0; i < b.mines; i++ {
		b.placeOneMine()
	}
}

func (b *Board) placeOneMine() {
	x := rand.Intn(b.width)
	y := rand.Intn(b.height)

	if b.Values[y][x] == 9 {
		b.placeOneMine()
	} else {
		b.Values[y][x] = 9
	}
}

func (b *Board) setNeighbors() {
	for y, j := range b.Values {
		for x, i := range j {
			if i != 9 {
				b.Values[y][x] = b.calculateMinesAround(x, y)
			}
		}
	}
}

func (b *Board) calculateMinesAround(x int, y int) int {
	initialX := func() int {
		if x == 0 {
			return 0
		} else {
			return x - 1
		}
	}()
	finalX := func() int {
		if x+1 > b.width-1 {
			return b.width - 1
		} else {
			return x + 1
		}
	}()
	initialY := func() int {
		if y == 0 {
			return 0
		} else {
			return y - 1
		}
	}()
	finalY := func() int {
		if y+1 > b.height-1 {
			return b.height - 1
		} else {
			return y + 1
		}
	}()

	counter := 0

	for j := initialY; j < finalY+1; j++ {
		for i := initialX; i < finalX+1; i++ {
			if b.Values[j][i] == 9 {
				counter += 1
			}
		}
	}

	return counter
}

func (b *Board) createCells() {
	for j, slice := range b.Values {
		for i, value := range slice {
			b.Cells[j][i] = cell.NewCell(value)
		}
	}
}

func NewBoard(width int, height int) *Board {
	values := make([][]int, height)
	cells := make([][]*cell.Cell, height)
	for i := range values {
		values[i] = make([]int, width)
		cells[i] = make([]*cell.Cell, width)
	}

	board := Board{
		Cells:  cells,
		Values: values,
		width:  width,
		height: height,
		mines:  99,
	}

	board.placeMines()
	board.setNeighbors()

	board.createCells()

	return &board
}
