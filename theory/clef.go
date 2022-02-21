package theory

type Clef struct {
	Type  ClefType
	Line  int
	Pitch Pitch
}

var ClefTreble Clef = Clef{
	Type:  ClefTypeG,
	Line:  2,
	Pitch: NewPitch(StepG, Natural, 4),
}

var ClefBass Clef = Clef{
	Type:  ClefTypeF,
	Line:  4,
	Pitch: NewPitch(StepF, Natural, 3),
}

var ClefAlto Clef = Clef{
	Type:  ClefTypeC,
	Line:  3,
	Pitch: MiddleC,
}

var ClefTenor Clef = Clef{
	Type:  ClefTypeC,
	Line:  4,
	Pitch: MiddleC,
}

type ClefType rune

const (
	ClefTypeG ClefType = 'G'
	ClefTypeF ClefType = 'F'
	ClefTypeC ClefType = 'C'
)
