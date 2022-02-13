package main

import (
	"errors"
	"trainer/glyphs"
	"trainer/notation"
	"trainer/notepaint"

	"github.com/veandco/go-sdl2/sdl"
)

type Program struct {
	SpaceHeight int
	Symbols     []notation.Symbol

	startPosition  sdl.Point
	fontSize       int
	painterBuilder *notepaint.Builder

	init         bool
	renderer     *sdl.Renderer
	shouldRender bool
}

func NewProgram() Program {
	startPosition := sdl.Point{
		X: 100,
		Y: 100,
	}
	fontSize := 90
	return Program{
		SpaceHeight:    10,
		Symbols:        []notation.Symbol{},
		startPosition:  startPosition,
		fontSize:       fontSize,
		painterBuilder: notepaint.NewBuilder(startPosition, fontSize),
	}
}

func (p *Program) Init(renderer *sdl.Renderer) error {
	if err := p.painterBuilder.Init(renderer); err != nil {
		return err
	}
	p.renderer = renderer
	p.init = true
	return nil
}

func (p *Program) clear() error {
	if err := p.renderer.SetDrawColor(255, 255, 255, 255); err != nil {
		return err
	}
	if err := p.renderer.Clear(); err != nil {
		return err
	}
	return nil
}

func (p *Program) Run() error {
	if !p.init {
		return errors.New("program not initialized")
	}
	if err := p.render(); err != nil {
		return err
	}
	for p.handleEvents() {
		if p.shouldRender {
			if err := p.render(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Program) render() error {
	if err := p.clear(); err != nil {
		return err
	}

	p.painterBuilder.Reset()
	if err := p.painterBuilder.AddSymbols(p.Symbols); err != nil {
		return err
	}

	notePainter := p.painterBuilder.Build()

	staffStore, err := notepaint.NewGlyphStore(p.fontSize, sdl.Color{A: 255}, p.renderer)
	if err != nil {
		return err
	}
	staffGlyph, err := staffStore.LoadGlyph(glyphs.Staff5LinesWide)
	if err != nil {
		return err
	}
	staffCount := float64(notePainter.CalculateWidth()) / float64(staffGlyph.Bounds().W)
	for i := 0; float64(i) < staffCount; i++ {
		err = staffGlyph.Render(p.renderer, sdl.Rect{
			X: p.startPosition.X + int32(i*int(staffGlyph.Bounds().W)),
			Y: p.startPosition.Y,
			W: staffGlyph.Bounds().W,
			H: staffGlyph.Bounds().H,
		})
		if err != nil {
			return err
		}
	}

	if err := notePainter.Paint(); err != nil {
		return err
	}

	p.renderer.Present()
	p.shouldRender = false
	return nil
}

func (p *Program) Free() {
	if p.painterBuilder != nil {
		p.painterBuilder.Free()
	}
}
