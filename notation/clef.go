package notation

type Clef struct {
	Type ClefType
	Line int
}
type ClefType int

const (
	ClefTypeNone = 0
	ClefTypeG    = 1
	ClefTypeF    = 2
	ClefTypeC    = 3
)
