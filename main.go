package main

import (
	"ai/ai"
	"ai/formula"
	"fmt"
)

func main() {

	x := [][]float64{{2.0}}
	y := []float64{3.0}

	// test1 := []float64{1.5, 2, 2.5, 3.0}
	// test2 := [][]float64{
	// 	{1.0, 0.75, 0.5, 0.25},
	// 	{0., 0.25, 0.5, 0.75},
	// }

	// fmt.Println(formula.MatrixMultiplication(test1, test2))

	w1 := [][]float64{{0.25, 0.5, 0.75, 1.0}}
	b1 := []float64{1.0, 1.0, 1.0, 1.0}

	w2 := [][]float64{
		{1.0, 0.75, 0.5, 0.25},
		{0., 0.25, 0.5, 0.75},
	}
	b2 := []float64{1.0, 1.0}

	fmt.Println(formula.MatrixMultiplication(x, w1))

	nn := ai.CreateNN()
	nn.SetWeights(w1, b1, 0)
	nn.SetWeights(w2, b2, 1)

	for _, wL := range nn.Layers {
		fmt.Println(wL.Weights)
	}

	nn.Train(x, y)
}
