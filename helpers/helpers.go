package helpers

func ExtractColumn(x [][]float64, rowNum int) []float64 {
	columns := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		columns[i] = x[i][rowNum]
	}

	return columns
}
