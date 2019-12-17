package textures

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var NumberMappings map[int]*sdl.Texture

// var ONE = createTexture("assets/numbers/1.bmp")
// var TWO = createTexture("assets/numbers/2.bmp")
// var THREE = createTexture("assets/numbers/3.bmp")
// var FOUR = createTexture("assets/numbers/4.bmp")
// var FIVE = createTexture("assets/numbers/5.bmp")
// var SIX = createTexture("assets/numbers/6.bmp")
// var SEVEN = createTexture("assets/numbers/7.bmp")
// var EIGHT = createTexture("assets/numbers/8.bmp")
// var NINE = createTexture("assets/numbers/9.bmp")

func CreateNumbers(renderer *sdl.Renderer) {
	NumberMappings = make(map[int]*sdl.Texture)
	for i := 0; i < 10; i++ {
		NumberMappings[i] = createTexture(fmt.Sprintf("assets/numbers/%v.bmp", i), renderer)
	}
}

func createTexture(path string, r *sdl.Renderer) *sdl.Texture {
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
