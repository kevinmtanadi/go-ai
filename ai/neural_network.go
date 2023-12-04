package ai

import (
	"fmt"
)

type neuralNetwork struct {
	Layers []Layer
}

type Layer struct {
	Weights [][]float64
	Bias    []float64
}

func CreateNN() *neuralNetwork {
	return &neuralNetwork{
		Layers: []Layer{},
	}
}

func (nn *neuralNetwork) SetWeights(w [][]float64, b []float64, nLayer int) *neuralNetwork {
	nn.Layers = append(nn.Layers, Layer{
		Weights: w,
		Bias:    b,
	})

	return nn
}

func (nn *neuralNetwork) Linear(inputDim int, outputDim int) *neuralNetwork {
	weightLayer := make([][]float64, outputDim)
	for i := 0; i < outputDim; i++ {
		weightLayer[i] = make([]float64, inputDim)
	}

	biasLayer := make([]float64, outputDim)

	newLayer := Layer{
		Weights: weightLayer,
		Bias:    biasLayer,
	}

	nn.Layers = append(nn.Layers, newLayer)

	return nn
}

func (nn *neuralNetwork) Train(x [][]float64, y []float64) {
	if len(x) != len(y) {
		panic(fmt.Sprintf("Dimension mismatch x and y {%d, %d}", len(x), len(y)))
	}

	for i := range x {
		// forward propagation
		outputLayers := [][]float64{}
		for l := range nn.Layers {
			if l == 0 {
				outputLayers = append(outputLayers, nn.forwardPropagation(x[i], l))
			} else {
				outputLayers = append(outputLayers, nn.forwardPropagation(outputLayers[l-1], l))
			}
		}

		// calculate loss

		// backward propagation
	}
}

func (nn *neuralNetwork) forwardPropagation(x []float64, nLayer int) []float64 {
	output := []float64{}

	var inputDim int
	var outputDim int

	if nLayer == 0 {
		inputDim = len(nn.Layers[nLayer].Weights)
	} else {
		inputDim = len(nn.Layers[nLayer].Weights[0])
	}
	outputDim = len(nn.Layers[nLayer].Bias)

	fmt.Println("inputDim: ", inputDim)
	fmt.Println("outputDim: ", outputDim)

	fmt.Println(output)
	return output
}

func (nn *neuralNetwork) backwardPropagation() {

}
