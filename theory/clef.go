package theory

type Clef struct {
	Type ClefType
	Line int
}

var ClefTreble Clef = Clef{
	Type: ClefTypeG,
	Line: 2,
}

var ClefBass Clef = Clef{
	Type: ClefTypeF,
	Line: 4,
}

var ClefAlto Clef = Clef{
	Type: ClefTypeC,
	Line: 3,
}

var ClefTenor Clef = Clef{
	Type: ClefTypeC,
	Line: 4,
}

type ClefType rune

const (
	ClefTypeG ClefType = 'G'
	ClefTypeF ClefType = 'F'
	ClefTypeC ClefType = 'C'
)
