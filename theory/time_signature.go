package theory

type TimeSignature struct {
	Beats int
	Unit  RhythmicValue
}

var Time44 TimeSignature = TimeSignature{
	Beats: 4,
	Unit:  Quarter,
}
