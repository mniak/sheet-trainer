package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (p *Program) handleEvents() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			if e.Repeat == 0 && e.State == sdl.PRESSED {
				if !p.handleKeyPress(e.Keysym) {
					return false
				}
			}
			break
		}
	}

	return true
}

func (p *Program) handleKeyPress(key sdl.Keysym) bool {
	fmt.Printf("Keyboard Event %d\n", key.Scancode)
	switch key.Scancode {
	case 82: // Key UP
		p.shouldRender = true
	}
	return true
}
