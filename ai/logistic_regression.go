package ai

import (
	"ai/formula"
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type LogisticRegression struct {
	W0           float64
	Wi           []float64
	Epochs       int
	LearningRate float64
	Cost         []float64
}

func (m *LogisticRegression) Train(x [][]float64, y []float64) {
	m.Wi = make([]float64, len(x[0]))
	m.W0 = 0.

	for epoch := 0; epoch < m.Epochs; epoch++ {
		yPredict := m.Predict(x)
		m.W0 = m.W0 - m.LearningRate*m.w0Exe(y, yPredict)
		for i := range m.Wi {
			m.Wi[i] = m.Wi[i] - m.LearningRate*m.wExe(x, y, yPredict, i)
		}
		costVal := m.costF(x, y, m.W0, m.Wi)
		m.Cost = append(m.Cost, costVal)
	}
}

func (m *LogisticRegression) Predict(x [][]float64) []float64 {
	yPredict := make([]float64, len(x))
	for i := range x {
		z := m.W0
		for j := range x[i] {
			z += m.Wi[j] * x[i][j]
		}

		yPredict[i] = 1.0 / (1.0 + formula.Exp(-z))
	}
	return yPredict
}

func (m *LogisticRegression) w0Exe(y, yPredict []float64) float64 {
	// return formula.Sum(formula.ArraySubtract(yPredict, y))

	sum := 0.

	for i := range y {
		sum += (yPredict[i] - y[i])
	}

	return sum
}

func (m *LogisticRegression) wExe(x [][]float64, y, yPredict []float64, featureIndex int) float64 {
	sum := 0.

	for i := range y {
		sum += (yPredict[i] - y[i]) * x[i][featureIndex]
	}

	return sum
	// return formula.Sum(formula.ArrayMultiplication(formula.ArraySubtract(yPredict, y), helpers.ExtractColumn(x, featureIndex)))
}

func (m *LogisticRegression) costF(x [][]float64, y []float64, w0 float64, wi []float64) float64 {
	yPredict := m.Predict(x)
	cost := 0.0

	for i := range x {
		cost += -y[i]*math.Log(yPredict[i]) - (1-y[i])*math.Log(1-yPredict[i])
	}

	return -cost
}

func (m *LogisticRegression) PlotLoss() {
	p := plot.New()

	points := make(plotter.XYs, len(m.Cost))
	for i, c := range m.Cost {
		points[i].X = float64(i + 1)
		points[i].Y = c
	}

	p.Title.Text = "Number of Epochs vs Loss"
	p.X.Label.Text = "Number of Epochs"
	p.Y.Label.Text = "Loss"

	s, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}

	p.Add(s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "loss.png"); err != nil {
		fmt.Println(err)
	}
}
