package preprocessing

import "ai/formula"

type Standardizer struct {
	StandardDeviation float64
	Mean              float64
}

// FitTransform
//
// Use this when declaring the standardizer.
// It will use the input data to standardize the future input
// when using Fit()
func (s *Standardizer) FitTransform(x []float64) []float64 {
	s.Mean = formula.Mean(x)
	s.StandardDeviation = formula.StandardDeviation(x, s.Mean)

	return s.Fit(x)
}

// Fit
//
// Standardize new data inputs
func (s *Standardizer) Fit(x []float64) []float64 {
	if s.Mean == 0 && s.StandardDeviation == 0 {
		panic("the standardizer hasn't been initialized yet. use FitTransform to initialize the standardizer")
	}
	standardizedData := make([]float64, len(x))
	for i, v := range x {
		standardizedData[i] = (v - s.Mean) / s.StandardDeviation
	}

	return standardizedData
}
