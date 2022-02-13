package main

import (
	"trainer/notation"
	"trainer/theory"
)

func StaffToSymbols(staff theory.Staff) []notation.Symbol {
	symbols := make([]notation.Symbol, 0)
	symbols = append(symbols, ClefToSymbol(staff.Clef))
	// symbols[1] = TimeSignatureToSymbol(
	for _, note := range staff.Notes {
		symbols = append(symbols, NoteToSymbol(staff.Clef, note))
	}
	return symbols
}

func ClefToSymbol(clef theory.Clef) notation.Symbol {
	var result notation.Clef
	switch clef.Type {
	case theory.ClefTypeG:
		result.Type = notation.ClefTypeG
	case theory.ClefTypeF:
		result.Type = notation.ClefTypeF
	case theory.ClefTypeC:
		result.Type = notation.ClefTypeC
	}

	result.Line = clef.Line
	return notation.Symbol{
		Clef: result,
	}
}

func NoteToSymbol(clef theory.Clef, note theory.Note) notation.Symbol {
	symbol := notation.Symbol{
		NoteHead: notation.NoteheadHalf,
	}

	symbol.Position = GetNotePosition(clef, note.Pitch)
	symbol.NoteHead = GetNoteHead(note.Duration)
	symbol.Stem = GetNoteStem(note.Duration, symbol.Position)
	symbol.Flag = GetNoteFlag(note.Duration)
	return symbol
}

func GetNotePosition(clef theory.Clef, pitch theory.Pitch) int {
	var position int
	switch clef.Type {
	case theory.ClefTypeF:
		position = int(clef.Line)*-2 - 10
		break
		// case theory.ClefTypeG:
		// 	break
		// case theory.ClefTypeC:
		// 	break
	}
	position += pitch.Octave * 7
	switch pitch.PitchClass {
	case theory.C:
		position += 0
		break
	case theory.D:
		position += 1
		break
	case theory.E:
		position += 2
		break
	case theory.F:
		position += 3
		break
	case theory.G:
		position += 4
		break
	case theory.A:
		position += 5
		break
	case theory.B:
		position += 6
		break
	}
	return position
}

func GetNoteHead(duration theory.RhythmicValue) notation.NoteHead {
	switch duration {
	case theory.Whole:
		return notation.NoteheadWhole
		break
	case theory.Half:
		return notation.NoteheadHalf
		break
	case theory.Quarter:
		return notation.NoteheadBlack
		break
	}
	return notation.NoteheadBlack
}

func GetNoteStem(duration theory.RhythmicValue, position int) notation.Stem {
	switch duration {
	case theory.Whole:
		return notation.StemNone
	}

	if position > 4 {
		return notation.StemDown
	}

	return notation.StemUp
}

func GetNoteFlag(duration theory.RhythmicValue) notation.Flag {
	switch duration {
	case theory.Eigth:
		return notation.FlagSingle
	case theory.Sixteenth:
		return notation.FlagDouble
	case theory.ThirtySecond:
		return notation.FlagTriple
	}
	return notation.FlagNone
}
