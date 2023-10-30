package preprocessing

import "ai/formula"

type Normalizer struct {
	Min float64
	Max float64
}

func (n *Normalizer) FitTransform(data [][]float64) [][]float64 {
	normalizedDatas := [][]float64{}

	for _, x := range data {
		n.Min = formula.Min(x)
		n.Max = formula.Max(x)

		normalizedDatas = append(normalizedDatas, n.Fit(x))
	}

	return normalizedDatas
}

func (n *Normalizer) FitMultiple(data [][]float64) [][]float64 {
	normalizedDatas := [][]float64{}

	for _, x := range data {
		normalizedDatas = append(normalizedDatas, n.Fit(x))
	}
	return normalizedDatas
}

func (n *Normalizer) Fit(x []float64) []float64 {
	if n.Min == 0 && n.Max == 0 {
		panic("the normalizer hasn't been initialized yet. use FitTransform to initialize the normalizer")
	}

	normalizedData := make([]float64, len(x))
	for i, v := range x {
		normalizedData[i] = (v - n.Min) / (n.Max - n.Min)
	}

	return normalizedData
}
