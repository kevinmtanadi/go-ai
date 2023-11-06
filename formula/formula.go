package formula

import (
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
		panic("Dimension mismatch")
	}

	var sum float64
	for i := 0; i < len(p1); i++ {
		sum += math.Pow(math.Abs(p1[i]-p2[i]), r)
	}

	return math.Pow(sum, 1/r)
}
