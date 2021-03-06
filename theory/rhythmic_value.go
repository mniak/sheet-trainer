package theory

type RhythmicValue uint16

const (
	Whole        RhythmicValue = 1
	Half         RhythmicValue = 2
	Quarter      RhythmicValue = 4
	Eigth        RhythmicValue = 8
	Sixteenth    RhythmicValue = 16
	ThirtySecond RhythmicValue = 32
)

func (v RhythmicValue) Value() float32 {
	return 1 / float32(v)
}
