package main

import (
	"trainer/notation"
	"trainer/theory"
)

func MeasuresToSymbols(measures []theory.Measure) []notation.Symbol {
	symbols := make([]notation.Symbol, 0)
	var previousClef theory.Clef
	for i, measure := range measures {
		if previousClef != measure.Clef {
			symbols = append(symbols, ClefToSymbol(measure.Clef))
			previousClef = measure.Clef
		}
		// if previousTimeSignature != measure.TimeSignature {
		// 	symbols[1] = TimeSignatureToSymbol(...)
		// 	previousTimeSignature = measure.TimeSignature
		// }
		for _, note := range measure.Notes {
			symbols = append(symbols, NoteToSymbol(measure.Clef, note))
		}
		isLast := i == len(measures)+1
		if !isLast {
			symbols = append(symbols, BarlineToSymbol())
		}
	}

	return symbols
}

func BarlineToSymbol() notation.Symbol {
	var result notation.Barline
	switch {
	default:
		result = notation.BarlineSimple
	}
	return notation.Symbol{
		Barline: result,
	}
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
		Note: notation.Note{
			NoteHead: GetNoteHead(note.Duration),
			Flag:     GetNoteFlag(note.Duration),
			Duration: note.Duration.Value(),
		},
		Position: GetNotePosition(clef, note.Pitch),
	}
	symbol.Note.Stem = GetNoteStem(note.Duration, symbol.Position)
	return symbol
}

func GetNotePosition(clef theory.Clef, pitch theory.Pitch) int {
	return (pitch.Octave()-clef.Pitch.Octave())*7 +
		int(pitch.Step()-clef.Pitch.Step()) +
		(clef.Line-1)*2
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
