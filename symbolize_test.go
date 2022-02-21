package main

import (
	"testing"
	"trainer/theory"

	"github.com/stretchr/testify/assert"
)

func TestGetNotePosition(t *testing.T) {
	testData := []struct {
		Name             string
		Clef             theory.Clef
		Pitch            theory.Pitch
		ExpectedPosition int
	}{
		{
			Name:             "G4 on treble L2",
			Clef:             theory.ClefTreble,
			Pitch:            theory.NewPitch(theory.StepG, theory.Natural, 4),
			ExpectedPosition: 2,
		},
		{
			Name:             "C4 on treble -L1",
			Clef:             theory.ClefTreble,
			Pitch:            theory.MiddleC,
			ExpectedPosition: -2,
		},
		{
			Name:             "C4 on alto L3",
			Clef:             theory.ClefAlto,
			Pitch:            theory.MiddleC,
			ExpectedPosition: 4,
		},
		{
			Name:             "C4 on tenor L4",
			Clef:             theory.ClefTenor,
			Pitch:            theory.MiddleC,
			ExpectedPosition: 6,
		},
		{
			Name:             "C4 on bass L6",
			Clef:             theory.ClefBass,
			Pitch:            theory.MiddleC,
			ExpectedPosition: 10,
		},
		{
			Name:             "D3 on bass L3",
			Clef:             theory.ClefBass,
			Pitch:            theory.NewPitch(theory.StepD, theory.Natural, 3),
			ExpectedPosition: 4,
		},
		{
			Name:             "C3 on bass S2",
			Clef:             theory.ClefBass,
			Pitch:            theory.NewPitch(theory.StepC, theory.Natural, 3),
			ExpectedPosition: 3,
		},
		{
			Name:             "B2 on bass L2",
			Clef:             theory.ClefBass,
			Pitch:            theory.NewPitch(theory.StepB, theory.Natural, 2),
			ExpectedPosition: 2,
		},
		{
			Name:             "B3 on clef -S2",
			Clef:             theory.ClefTreble,
			Pitch:            theory.NewPitch(theory.StepB, theory.Natural, 3),
			ExpectedPosition: -3,
		},
	}

	for _, td := range testData {
		t.Run(td.Name, func(t *testing.T) {
			actual := GetNotePosition(td.Clef, td.Pitch)
			assert.Equal(t, td.ExpectedPosition, actual)
		})
	}
}
