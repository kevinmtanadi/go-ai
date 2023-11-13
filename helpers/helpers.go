package helpers

func ExtractColumn(x [][]float64, rowNum int) []float64 {
	columns := make([]float64, len(x))

	for i := 0; i < len(x[0]); i++ {
		columns[i] = x[rowNum][i]
	}

	return columns
}
