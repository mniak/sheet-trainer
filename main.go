package main

import (
	"log"
	"trainer/generators"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
		// panic(err)
	}
}

func main() {
	handle(sdl.Init(sdl.INIT_EVERYTHING))
	defer sdl.Quit()

	handle(ttf.Init())
	defer ttf.Quit()

	window, err := sdl.CreateWindow("test",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1920, 1080,
		sdl.WINDOW_SHOWN|sdl.WINDOW_FULLSCREEN)
	handle(err)
	defer window.Destroy()

	surf, err := window.GetSurface()
	handle(err)
	defer surf.Free()

	renderer, err := window.GetRenderer()
	handle(err)
	defer renderer.Destroy()

	program := NewProgram()
	defer program.Free()

	measures := generators.ToArray(generators.Exercise1)
	symbols := MeasuresToSymbols(measures)
	program.Symbols = symbols
	handle(program.Init(renderer))
	handle(program.Run())
}
