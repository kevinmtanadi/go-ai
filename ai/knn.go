package ai

import (
	"ai/formula"
	"fmt"
	"sync"
)

type kNN struct {
	k           int
	p           int
	initialData [][]float64
	outputData  []float64
}

func NewKNN(k int, p int, initialData [][]float64, outputData []float64) *kNN {
	return &kNN{
		k:           k,
		p:           p,
		initialData: initialData,
		outputData:  outputData,
	}
}

func (knn *kNN) GoPredict(input [][]float64) []float64 {
	var output []float64

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, data := range input {
		wg.Add(1)

		go func(data []float64) {
			defer wg.Done()

			var distances [][]float64

			for i, d := range knn.initialData {
				distances = append(distances, []float64{formula.MinkowskiDistance(d, data, float64(knn.p)), knn.outputData[i]})
			}

			formula.Sort(distances)

			var neighbors []float64

			for i := 0; i < knn.k; i++ {
				neighbors = append(neighbors, distances[i][1])
			}

			prediction := MostCommon(neighbors)

			mu.Lock()
			output = append(output, prediction)
			mu.Unlock()

		}(data)
	}

	wg.Wait()

	return output
}

func (knn *kNN) Predict(input [][]float64) []float64 {
	var output []float64

	for _, data := range input {
		var distances [][]float64

		for i, d := range knn.initialData {
			if len(d) != len(data) {
				fmt.Println("something wrong on: ", i)
			}
			distances = append(distances, []float64{formula.MinkowskiDistance(d, data, float64(knn.p)), knn.outputData[i]})
		}

		formula.Sort(distances)

		var neighbors []float64

		for i := 0; i < knn.k; i++ {
			neighbors = append(neighbors, distances[i][1])
		}

		prediction := MostCommon(neighbors)

		output = append(output, prediction)

	}

	return output
}

func MostCommon(data []float64) float64 {
	counts := make(map[float64]int)

	for _, d := range data {
		counts[d]++
	}

	max := 0
	neighbor := 0.

	for k, v := range counts {
		if v > max {
			max = v
			neighbor = k
		}
	}

	return neighbor
}
