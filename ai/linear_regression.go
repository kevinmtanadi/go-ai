package ai

import (
	"ai/formula"
	"ai/helpers"
	"math"
)

type LinearRegression struct {
	W0           float64
	Wi           []float64
	Epochs       int
	LearningRate float64
	Cost         []float64
}

func (m *LinearRegression) Train() {

}

func (m *LinearRegression) w0Exe(y, yPredict []float64) float64 {
	n := float64(len(y))
	ans := formula.Sum(formula.ArraySubtract(yPredict, y))

	return (math.Pow(ans, -2) / n)
}

func (m *LinearRegression) wExe(x [][]float64, y, yPredict []float64, featureIndex int) float64 {
	n := float64(len(y))
	ans := formula.Sum(formula.ArrayMultiplication(formula.ArraySubtract(yPredict, y), helpers.ExtractColumn(x, featureIndex)))
	return (math.Pow(ans, -2) / n) * ans
}

// func (m *LinearRegression) yExe(x [][]float64) float64 {
// 	return m.W0 + formula.Dot(x, m.Wi)
// }
