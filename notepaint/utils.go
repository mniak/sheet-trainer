package notepaint

import (
	"trainer/glyphs"
	"trainer/notation"
)

func GetNoteHeadGlyph(head notation.NoteHead) rune {
	switch head {
	case notation.NoteheadWhole:
		return glyphs.NoteheadWhole
	case notation.NoteheadHalf:
		return glyphs.NoteheadHalf
	}
	return glyphs.NoteheadBlack
}

func GetClefGlyph(clefType notation.ClefType) rune {
	switch clefType {
	case notation.ClefTypeF:
		return glyphs.ClefF
	case notation.ClefTypeC:
		return glyphs.ClefC
	}
	return glyphs.ClefG
}
