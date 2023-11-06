package ai

import (
	"ai/formula"
)

type kNNRegression struct {
	k           int
	p           int
	initialData [][]float64
	outputData  []float64
}

func NewKNNRegression(k int, p int, initialData [][]float64, outputData []float64) *kNNRegression {
	return &kNNRegression{
		k:           k,
		p:           p,
		initialData: initialData,
		outputData:  outputData,
	}
}

func (knn *kNNRegression) Predict(input [][]float64) []float64 {
	var output []float64

	for _, data := range input {
		var distances [][]float64

		for i, d := range knn.initialData {
			distances = append(distances, []float64{formula.MinkowskiDistance(d, data, float64(knn.p)), knn.outputData[i]})
		}

		formula.Sort(distances)

		var neighbors []float64

		for i := 0; i < knn.k; i++ {
			neighbors = append(neighbors, distances[i][1])
		}

		predicition := formula.Mean(neighbors)

		output = append(output, predicition)
	}

	return output
}
