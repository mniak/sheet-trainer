package theory

type Note struct {
	Pitch    Pitch
	Duration RhythmicValue
	Rest     bool
}

func Rest(value RhythmicValue) Note {
	return Note{
		Rest:     true,
		Duration: value,
	}
}
