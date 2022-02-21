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

func GetBarlineGlyph(barline notation.Barline) rune {
	switch barline {
	case notation.BarlineSimple:
		return glyphs.BarlineSimple
	case notation.BarlineDouble:
		return glyphs.BarlineDouble
	case notation.BarlineFinal:
		return glyphs.BarlineFinal
	case notation.BarlineRepeatBegin:
		return glyphs.BarlineRepeatBegin
	case notation.BarlineRepeatEnd:
		return glyphs.BarlineRepeatEnd
	case notation.BarlineRepeatBeginEnd:
		return glyphs.BarlineRepeatBeginEnd
	}
	return glyphs.ClefG
}
