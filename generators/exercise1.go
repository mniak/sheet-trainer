package generators

import "trainer/theory"

func Exercise1(cgen chan theory.Measure) {
	for i := 0; i < 7; i++ {
		for _, m := range exercise1_line(i) {
			cgen <- m
		}
	}
}

func exercise1_line(n int) []theory.Measure {
	result := make([]theory.Measure, 0)

	pitch := theory.NewPitch(theory.StepD, theory.Natural, 3).UpN(n)
	result = append(result, theory.Measure{
		Clef:          theory.ClefBass,
		TimeSignature: theory.Time44,
		Notes: []theory.Note{
			{
				Pitch:    pitch,
				Duration: theory.Whole,
			},
		},
	})
	result = append(result, theory.Measure{
		Clef:          theory.ClefBass,
		TimeSignature: theory.Time44,
		Notes: []theory.Note{
			{
				Pitch:    pitch,
				Duration: theory.Half,
			},
			{
				Pitch:    pitch,
				Duration: theory.Half,
			},
		},
	})
	result = append(result, theory.Measure{
		Clef:          theory.ClefBass,
		TimeSignature: theory.Time44,
		Notes: []theory.Note{
			{
				Pitch:    pitch,
				Duration: theory.Quarter,
			},
			{
				Pitch:    pitch,
				Duration: theory.Quarter,
			},
			{
				Pitch:    pitch,
				Duration: theory.Quarter,
			},
			{
				Pitch:    pitch,
				Duration: theory.Quarter,
			},
		},
	})
	result = append(result, theory.Measure{
		Clef:          theory.ClefBass,
		TimeSignature: theory.Time44,
		Notes: []theory.Note{
			{
				Pitch:    pitch,
				Duration: theory.Half,
			},
			{
				Pitch:    pitch,
				Duration: theory.Quarter,
			},
			{
				Pitch:    pitch,
				Duration: theory.Quarter,
			},
		},
	})

	return result
}
