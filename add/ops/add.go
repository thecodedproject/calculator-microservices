package ops

func Add(inputs []float64) float64 {
	if len(inputs) == 0 {
		return 0.0
	}

	var sum float64

	for _, v := range inputs {
		sum += v
	}

	return sum
}
