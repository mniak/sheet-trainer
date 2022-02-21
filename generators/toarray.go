package generators

import "trainer/theory"

type GeneratorFunc func(cgen chan theory.Measure)

func ToArray(generate GeneratorFunc) []theory.Measure {
	ch := make(chan theory.Measure)
	defer close(ch)
	result := make([]theory.Measure, 0)
	go func() {
		for m := range ch {
			result = append(result, m)
		}
	}()
	generate(ch)
	return result
}
