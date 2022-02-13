package theory

type Staff struct {
	Clef          Clef
	TimeSignature TimeSignature
	Notes         []Note
}

type TimeSignature struct {
	Beats int
	Unit  RhythmicValue
}

var Time44 TimeSignature = TimeSignature{
	Beats: 4,
	Unit:  Quarter,
}
