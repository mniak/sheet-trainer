package notation

type NoteHead rune

const (
	NoteheadWhole NoteHead = 1
	NoteheadHalf  NoteHead = 2
	NoteheadBlack NoteHead = 3
)

type Symbol struct {
	Position int
	Clef     Clef
	Note     Note
	Bar      Bar
}

type SymbolType string

const (
	SymbolTypeClef SymbolType = "clef"
	SymbolTypeNote SymbolType = "note"
)

func (s *Symbol) Type() SymbolType {
	if s.Clef.Type != ClefTypeNone {
		return SymbolTypeClef
	}

	return SymbolTypeNote
}
