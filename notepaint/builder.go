package notepaint

import (
	"errors"
	"trainer/glyphs"
	"trainer/notation"

	"github.com/veandco/go-sdl2/sdl"
)

type Builder struct {
	initialOffset sdl.Point
	instructions  []PaintInstructions
	fontSize      int

	init   bool
	store  *GlyphStore
	offset sdl.Point
	hdpi   float32
}

func NewBuilder(startingPoint sdl.Point, fontSize int) *Builder {
	return &Builder{
		fontSize:      fontSize,
		initialOffset: startingPoint,
		instructions:  make([]PaintInstructions, 0),
	}
}

func (p *Builder) Init(renderer *sdl.Renderer) error {
	noteColor := sdl.Color{
		R: 127,
		G: 127,
		B: 255,
		A: 255,
	}

	var err error
	p.store, err = NewGlyphStore(p.fontSize, noteColor, renderer)
	if err != nil {
		return err
	}

	p.init = true
	_, p.hdpi, _, err = sdl.GetDisplayDPI(0)
	return nil
}

func (p *Builder) drawTexture(glyph *Glyph) {
	dst := sdl.Rect{
		X: p.offset.X,
		Y: p.offset.Y,
		W: glyph.Bounds().W,
		H: glyph.Bounds().H,
	}
	p.instructions = append(p.instructions, PaintInstructions{
		Glyph:    glyph,
		Position: dst,
	})
}

func (p *Builder) addClef(symbol notation.Symbol) error {
	p.moveUp((symbol.Clef.Line - 1) * 2)
	char := GetClefGlyph(symbol.Clef.Type)

	p.moveRight(20)

	clef, err := p.store.Draw(p, char)
	if err != nil {
		return err
	}

	p.moveRightAbsolute(clef.W)
	p.moveRightAbsolute(20)

	return nil
}

func (p *Builder) addNote(symbol notation.Symbol) error {
	p.moveUp(symbol.Position)

	noteheadGlyph := GetNoteHeadGlyph(symbol.Note.NoteHead)
	head, err := p.store.Draw(p, noteheadGlyph)
	if err != nil {
		return err
	}

	p.moveRightAbsolute(head.W)
	defer p.moveRight(20)
	if symbol.Note.Stem == notation.StemNone {
		return nil
	}

	stem, err := p.store.LoadGlyph(glyphs.Stem)
	if err != nil {
		return err
	}

	p.moveRightAbsolute(-stem.Bounds().W)
	p.drawTexture(stem)
	p.moveRightAbsolute(stem.bounds.W / 2)
	defer p.moveRightAbsolute(stem.bounds.W - (stem.bounds.W / 2))

	if symbol.Note.Flag == notation.FlagNone {
		return nil
	}

	p.moveUp(7)
	_, err = p.store.Draw(p, glyphs.Flag8Up)
	if err != nil {
		return err
	}

	return nil
}

func (p *Builder) addBarline(symbol notation.Symbol) error {
	glyph := GetBarlineGlyph(symbol.Barline)
	box, err := p.store.Draw(p, glyph)
	if err != nil {
		return err
	}
	p.moveRightAbsolute(box.W)
	p.moveRight(1)
	return nil
}

func (p *Builder) moveUp(steps int) {
	p.offset.Y -= int32(float32(steps) * float32(p.fontSize) / 8)
}

func (p *Builder) moveRightAbsolute(units int32) {
	p.offset.X += units
}

func (p *Builder) moveRight(units float32) {
	p.offset.X += int32(float32(units) * float32(p.fontSize) / p.hdpi)
}

func (p *Builder) AddSymbols(symbols []notation.Symbol) error {
	if !p.init {
		return errors.New("painter was not initialized")
	}

	for _, s := range symbols {
		originalY := p.offset.Y
		switch s.Type() {
		case notation.SymbolTypeClef:
			if err := p.addClef(s); err != nil {
				return err
			}
		case notation.SymbolTypeBarline:
			if err := p.addBarline(s); err != nil {
				return err
			}
		default:
			if err := p.addNote(s); err != nil {
				return err
			}
		}
		p.offset.Y = originalY
	}
	return nil
}

func (p *Builder) Reset() {
	p.offset = p.initialOffset
}

func (p *Builder) Free() {
	if p.store != nil {
		p.store.Free()
	}
}
