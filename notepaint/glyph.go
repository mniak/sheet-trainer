package notepaint

import "github.com/veandco/go-sdl2/sdl"

type Glyph struct {
	texture *sdl.Texture
	bounds  sdl.Rect
}

func NewGlyph(texture *sdl.Texture, bounds sdl.Rect) *Glyph {
	return &Glyph{
		texture: texture,
		bounds:  bounds,
	}
}

func (g *Glyph) Destroy() error {
	if g.texture != nil {
		return g.texture.Destroy()
	}
	return nil
}

func (g *Glyph) Bounds() sdl.Rect {
	return g.bounds
}

func (g *Glyph) Render(renderer *sdl.Renderer, dest sdl.Rect) error {
	return renderer.Copy(g.texture, &g.bounds, &dest)
}
