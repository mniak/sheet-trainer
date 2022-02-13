package notation

type Stem int

const (
	StemNone Stem = 0
	StemUp   Stem = 1
	StemDown Stem = -1
)

type Flag string

const (
	FlagNone Flag = ""

	FlagSingle      Flag = "1"
	FlagSingleIn    Flag = "-1"
	FlagSingleOut   Flag = "1-"
	FlagSingleInOut Flag = "-1-"

	FlagDouble      Flag = "2"
	FlagDoubleIn    Flag = "-2"
	FlagDoubleOut   Flag = "2-"
	FlagDoubleInOut Flag = "-2-"

	FlagTriple      Flag = "3"
	FlagTripleIn    Flag = "-3"
	FlagTripleOut   Flag = "3-"
	FlagTripleInOut Flag = "-3-"
)
