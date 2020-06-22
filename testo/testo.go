package testo

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Summarized(xs []float64) float64 {
	total := float64(0)
	for i := 0; i < len(xs); i++ {
		total += xs[i]
	}
	return total
}
