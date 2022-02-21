package notation

type Clef struct {
	Type ClefType
	Line int
}
type ClefType int

const (
	ClefTypeNone ClefType = iota
	ClefTypeG
	ClefTypeF
	ClefTypeC
)
