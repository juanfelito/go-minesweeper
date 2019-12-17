package main

import (
	"fmt"
	"go-minesweeper/pkg/board"
	"go-minesweeper/pkg/textures"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WIDTH  = 30
	HEIGHT = 18
)

func main() {
	numberMatrix := board.NewBoard(WIDTH, HEIGHT)
	numberMatrix.PrintBoard()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Felipe's Minesweeper",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WIDTH*40, HEIGHT*40,
		sdl.WINDOW_OPENGL,
	)

	if err != nil {
		fmt.Println("Initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer:", err)
	}
	defer renderer.Destroy()

	textures.CreateNumbers(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for j, slice := range numberMatrix.Values {
			for i, value := range slice {
				renderer.Copy(textures.NumberMappings[value], &sdl.Rect{X: 0, Y: 0, W: 20, H: 20}, &sdl.Rect{X: int32(i * 40), Y: int32(j * 40), W: 40, H: 40})
			}
		}

		renderer.Present()
	}
}
