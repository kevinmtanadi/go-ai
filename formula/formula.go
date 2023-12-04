package formula

import (
	"fmt"
	"math"
	"sort"
)

func Dot(x, y []float64) float64 {
	result := 0.0
	for i := 0; i < len(x); i++ {
		result += x[i] * y[i]
	}

	return result
}

func Exp(x float64) float64 {
	return math.Pow(2.718281828, x)
}

func Sum(x []float64) float64 {
	result := 0.0
	for i := 0; i < len(x); i++ {
		result += x[i]
	}

	return result
}

func Sigmoid(x float64) float64 {
	return 1 / (1 + Exp(-x))
}

// ArraySubtract:
//
// Returns subtraction between two arrays with same length
func ArraySubtract(x, y []float64) []float64 {
	if len(x) != len(y) {
		panic("ArraySubtract: len(x) != len(y)")
	}

	newArr := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		newArr[i] = x[i] - y[i]
	}

	return newArr
}

func ArrayMultiplication(x, y []float64) []float64 {
	if len(x) != len(y) {
		panic("ArrayMultiplication: len(x) != len(y)")
	}

	result := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		result[i] = x[i] * y[i]
	}

	return result
}

func Mean(x []float64) float64 {
	return Sum(x) / float64(len(x))
}

func StandardDeviation(x []float64, mean float64) float64 {
	std := 0.

	for _, v := range x {
		std += math.Pow(v-mean, 2)
	}

	return math.Sqrt(std / float64(len(x)))
}

func Min(x []float64) float64 {
	min := math.Inf(1)
	for _, v := range x {
		if v < min {
			min = v
		}
	}

	return min
}

func Max(x []float64) float64 {
	max := math.Inf(-1)
	for _, v := range x {
		if v > max {
			max = v
		}
	}

	return max
}

func Correlation(x, y []float64) float64 {
	n := float64(len(x))

	pembilang := (n*Dot(x, y) - (Sum(x) * Sum(y)))
	pembagi1 := n*Sum(ArrayMultiplication(x, x)) - math.Pow(Sum(x), 2)
	pembagi2 := n*Sum(ArrayMultiplication(y, y)) - math.Pow(Sum(y), 2)

	corr := pembilang / math.Sqrt(pembagi1*pembagi2)
	return corr
}

func MergeSort(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []float64) []float64 {
	result := make([]float64, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func Sort(data [][]float64) {
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
}

func MinkowskiDistance(p1, p2 []float64, r float64) float64 {
	if len(p1) != len(p2) {
		panic(fmt.Sprintf("Dimension mismatch {%d, %d}", len(p1), len(p2)))
	}

	var sum float64
	for i := 0; i < len(p1); i++ {
		sum += math.Pow(math.Abs(p1[i]-p2[i]), r)
	}

	return math.Pow(sum, 1/r)
}

func MatrixMultiplication(x [][]float64, y [][]float64) [][]float64 {
	r1 := len(x)
	c1 := len(x[0])
	r2 := len(y)
	c2 := len(y[0])

	if c1 != r2 {
		panic(fmt.Sprintf("Dimension mismatch {%d, %d} x {%d, %d}", r1, c1, r2, c2))
	}

	result := make([][]float64, r1)
	for i := range result {
		result[i] = make([]float64, c2)
	}

	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			for k := 0; k < c1; k++ {
				result[i][j] += x[i][k] * y[k][j]
			}
		}
	}

	return result
}

func Transpose(x [][]float64) [][]float64 {
	result := [][]float64{}
	for i := 0; i < len(x[0]); i++ {
		row := []float64{}
		for j := 0; j < len(x); j++ {
			row = append(row, x[j][i])
		}
		result = append(result, row)
	}

	return result
}
