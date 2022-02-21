package notation

type Barline int

const (
	BarlineNone Barline = iota
	BarlineSimple
	BarlineDouble
	BarlineFinal
	BarlineRepeatBegin
	BarlineRepeatEnd
	BarlineRepeatBeginEnd
)
