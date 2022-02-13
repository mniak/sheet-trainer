package main

import (
	"log"
	"trainer/theory"

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
		800, 600,
		sdl.WINDOW_SHOWN)
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

	notes := GenerateNotes()
	clef := theory.ClefBass
	staff := theory.Staff{
		Clef:  clef,
		Notes: notes,
	}
	symbols := StaffToSymbols(staff)
	program.Symbols = symbols
	handle(program.Init(renderer))
	handle(program.Run())
}

func GenerateNotes() []theory.Note {
	pitch := theory.Pitch{
		PitchClass: theory.D,
		Octave:     3,
	}
	return []theory.Note{
		{
			Pitch:    pitch,
			Duration: theory.Whole,
		},

		{
			Pitch:    pitch,
			Duration: theory.Half,
		},
		{
			Pitch:    pitch,
			Duration: theory.Half,
		},

		{
			Pitch:    pitch,
			Duration: theory.Quarter,
		},
		{
			Pitch:    pitch,
			Duration: theory.Quarter,
		},
		{
			Pitch:    pitch,
			Duration: theory.Quarter,
		},
		{
			Pitch:    pitch,
			Duration: theory.Quarter,
		},

		{
			Pitch:    pitch,
			Duration: theory.Half,
		},
		{
			Pitch:    pitch,
			Duration: theory.Quarter,
		},
		{
			Pitch: theory.Pitch{
				PitchClass: theory.B,
				Octave:     2,
			},
			Duration: theory.Sixteenth,
		},
	}
}
