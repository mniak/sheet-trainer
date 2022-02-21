package theory

func NewPitch(step PitchStep, alteration PitchAlteration, octave uint8) Pitch {
	return Pitch{
		step:       step,
		alteration: alteration,
		octave:     octave,
	}
}

type Pitch struct {
	step       PitchStep
	alteration PitchAlteration
	octave     uint8
}

func (p *Pitch) Step() PitchStep {
	return p.step
}

func (p *Pitch) Alteration() PitchAlteration {
	return p.alteration
}

func (p *Pitch) Octave() int {
	return int(p.octave)
}

var MiddleC = Pitch{
	step:   StepC,
	octave: 4,
}

type PitchStep int

const (
	StepC PitchStep = 0
	StepD PitchStep = 1
	StepE PitchStep = 2
	StepF PitchStep = 3
	StepG PitchStep = 4
	StepA PitchStep = 5
	StepB PitchStep = 6
)

func (p PitchStep) Up() PitchStep {
	return PitchStep(p + 1%7)
}

func (p PitchStep) Down() PitchStep {
	return PitchStep(p - 1%7)
}

type PitchAlteration int

const (
	DoubleFlat  PitchAlteration = -2
	Flat        PitchAlteration = -1
	Natural     PitchAlteration = 0
	Sharp       PitchAlteration = 1
	DoubleSharp PitchAlteration = 2
)

func (p Pitch) OctaveUp() Pitch {
	p.octave++
	return p
}

func (p Pitch) OctaveDown() Pitch {
	p.octave--
	return p
}

func (p Pitch) Sharpen() Pitch {
	p.alteration++
	return p
}

func (p Pitch) Flatten() Pitch {
	p.alteration--
	return p
}

func (p Pitch) Up() Pitch {
	p.step = p.step.Up()
	if p.step == StepC {
		return p.OctaveUp()
	}
	return p
}

func (p Pitch) Down() Pitch {
	p.step = p.step.Down()
	if p.step == StepB {
		return p.OctaveDown()
	}
	return p
}

func (p Pitch) UpN(n int) Pitch {
	for i := 0; i < n; i++ {
		p = p.Up()
	}
	return p
}

func (p Pitch) DownN(n int) Pitch {
	for i := 0; i < n; i++ {
		p = p.Down()
	}
	return p
}
