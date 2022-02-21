package notepaint

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type GlyphStore struct {
	font         *ttf.Font
	color        sdl.Color
	loadedGlyphs map[rune]*Glyph
	renderer     *sdl.Renderer
}

func NewGlyphStore(fontSize int, color sdl.Color, renderer *sdl.Renderer) (*GlyphStore, error) {
	font, err := ttf.OpenFont("./fonts/Bravura.otf", fontSize)
	if err != nil {
		return nil, err
	}
	return &GlyphStore{
		font:         font,
		color:        color,
		loadedGlyphs: make(map[rune]*Glyph),
		renderer:     renderer,
	}, nil
}

func (g *GlyphStore) LoadGlyph(r rune) (*Glyph, error) {
	if tex, ok := g.loadedGlyphs[r]; ok {
		return tex, nil
	}
	surf, err := g.font.RenderGlyphBlended(r, g.color)
	if err != nil {
		return nil, err
	}
	defer surf.Free()
	tex, err := g.renderer.CreateTextureFromSurface(surf)
	glyph := NewGlyph(tex, surf.ClipRect)
	g.loadedGlyphs[r] = glyph
	return glyph, nil
}

func (g *GlyphStore) Draw(painter *Builder, r rune) (sdl.Rect, error) {
	glyph, err := g.LoadGlyph(r)
	if err != nil {
		return sdl.Rect{}, err
	}
	painter.drawTexture(glyph)
	return glyph.Bounds(), nil
}

func (g *GlyphStore) Free() {
	if g.font != nil {
		g.font.Close()
	}
	for _, tex := range g.loadedGlyphs {
		tex.Destroy()
	}
}
