package textures

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var NumberMappings map[int]*sdl.Texture

func CreateNumbers(renderer *sdl.Renderer) {
	NumberMappings = make(map[int]*sdl.Texture)
	for i := 0; i < 10; i++ {
		NumberMappings[i] = CreateTexture(fmt.Sprintf("assets/numbers/%v.bmp", i), renderer)
	}
}

func CreateTexture(path string, r *sdl.Renderer) *sdl.Texture {
	img, loadErr := sdl.LoadBMP(path)
	if loadErr != nil {
		fmt.Println("Loading button:", loadErr)
		return nil
	}

	texture, err := r.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("Creating texture:", err)
		return nil
	}

	return texture
}
