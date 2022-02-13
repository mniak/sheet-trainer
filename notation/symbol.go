package notation

type NoteHead rune

const (
	NoteheadWhole NoteHead = 1
	NoteheadHalf  NoteHead = 2
	NoteheadBlack NoteHead = 3
)

type Symbol struct {
	Clef     Clef
	NoteHead NoteHead
	Position int
	Stem     Stem
	Flag     Flag
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
