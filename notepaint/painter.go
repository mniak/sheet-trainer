package notepaint

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type PaintInstructions struct {
	Glyph    *Glyph
	Position sdl.Rect
}

type NotePainter struct {
	instructions []PaintInstructions
	renderer     *sdl.Renderer
}

func (b *Builder) Build() *NotePainter {
	return &NotePainter{
		instructions: b.instructions,
		renderer:     b.store.renderer,
	}
}

func (p *NotePainter) CalculateWidth() int {
	var max float64 = 0
	for _, instr := range p.instructions {
		max = math.Max(max, float64(instr.Position.X+instr.Position.W))
	}
	return int(max)
}

func (p *NotePainter) Paint() error {
	for _, instr := range p.instructions {
		p.renderer.Copy(instr.Glyph.texture, &instr.Glyph.bounds, &instr.Position)
	}
	return nil
}
