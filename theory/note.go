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

type Pitch struct {
	PitchClass PitchClass
	Octave     int
}

func (p Pitch) OctaveUp() Pitch {
	p.Octave++
	return p
}

func (p Pitch) OctaveDown() Pitch {
	p.Octave--
	return p
}

type PitchClass string

var MiddleC = Pitch{
	PitchClass: C,
	Octave:     4,
}

const (
	CFlat  PitchClass = "Cb"
	C      PitchClass = "C"
	CSharp PitchClass = "C#"

	DFlat  PitchClass = "Db"
	D      PitchClass = "D"
	DSharp PitchClass = "D#"

	EFlat  PitchClass = "Eb"
	E      PitchClass = "E"
	ESharp PitchClass = "E#"

	FFlat  PitchClass = "Fb"
	F      PitchClass = "F"
	FSharp PitchClass = "F#"

	GFlat  PitchClass = "Gb"
	G      PitchClass = "G"
	GSharp PitchClass = "G#"

	AFlat  PitchClass = "Ab"
	A      PitchClass = "A"
	ASharp PitchClass = "A#"

	BFlat  PitchClass = "Bb"
	B      PitchClass = "B"
	BSharp PitchClass = "B#"
)
