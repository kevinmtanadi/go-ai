package ai

import (
	"ai/formula"
	"ai/helpers"
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
		yPredict := m.Predict(x, m.W0, m.Wi)
		m.W0 = m.W0 - m.LearningRate*m.w0Exe(y, yPredict)
		wexe := [][]float64{}
		for i := 0; i < len(x[0]); i++ {
			wexe = append(wexe, m.wExe(x, y, yPredict, i))
		}

		wExes := sum2DArray(wexe)
		fmt.Println("len wexes", len(wExes))
		for i, w := range wExes {
			m.Wi[i] = m.Wi[i] - m.LearningRate*w
		}
		// for i := range m.Wi {
		// 	m.Wi[i] = m.Wi[i] - m.LearningRate*m.wExe(x, y, yPredict, i)
		// }
		costVal := m.costF(x, y, m.W0, m.Wi)
		m.Cost = append(m.Cost, costVal)
	}
}

func (m *LogisticRegression) Predict(x [][]float64, w0 float64, w1 []float64) []float64 {
	var predictions []float64
	for _, row := range x {
		z := w0
		for i := range row {
			z += row[i] * w1[i]
		}
		predictions = append(predictions, Sigmoid(z))
	}
	return predictions
}

func Sigmoid(z float64) float64 {
	return 1.0 / (1.0 + formula.Exp(-z))
}

func (m *LogisticRegression) w0Exe(y, yPredict []float64) float64 {
	sum := 0.
	for i := range y {
		sum += yPredict[i] - y[i]
	}

	return sum
}

func (m *LogisticRegression) wExe(x [][]float64, y, yPredict []float64, featureIndex int) []float64 {
	return formula.ArrayMultiplication(formula.ArraySubtract(yPredict, y), helpers.ExtractColumn(x, featureIndex))
}

func (m *LogisticRegression) costF(x [][]float64, y []float64, w0 float64, w1 []float64) float64 {
	yPredict := m.Predict(x, w0, w1)
	n := len(y)
	cost := 0.0
	for i := 0; i < n; i++ {
		cost += y[i]*math.Log(yPredict[i]) + (1-y[i])*math.Log(1-yPredict[i])
	}
	return -cost
}

// func (m *LogisticRegression) Test() {
// 	x := [][]float64{
// 		{0.1, 0.2},
// 		{0.1, 0.2},
// 	}

// 	y := []float64{0, 0}
// 	yPred := []float64{0, 1}
// 	// w0 := 32.411240798068064
// 	// w1 := []float64{-24.87726705294277, -34.99987387351453}

// 	wexe := [][]float64{}
// 	for i := 0; i < len(x); i++ {
// 		wexe = append(wexe, m.wExe(x, y, yPred, i))
// 	}

// 	fmt.Println(sum2DArray(wexe))
// }

func sum2DArray(x [][]float64) []float64 {
	rows := len(x)
	cols := len(x[0])

	result := make([]float64, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j] += x[i][j]
		}
	}

	return result
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
