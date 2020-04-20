package cell

import (
	"go-minesweeper/pkg/textures"

	"github.com/veandco/go-sdl2/sdl"
)

type Cell struct {
	Texture *sdl.Texture
}

func NewCell(aob int) *Cell {
	return &Cell{
		Texture: textures.NumberMappings[aob],
	}
}
