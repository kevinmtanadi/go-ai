package main

import (
	"ai/dataframe"
)

func main() {
	// data := []float64{1.1, 0, 100.0, 34.2, 3.14, 2.718, 96.23}

	// standardizer := preprocessing.Standardizer{}
	// standardizedData := standardizer.FitTransform(data)
	// fmt.Println(standardizedData)

	// normalizer := preprocessing.Normalizer{}
	// normalizedData := normalizer.FitTransform(data)
	// fmt.Println(normalizedData)

	df := dataframe.DataFrame{}
	df.ReadCSV("data.csv")

	df.OneHotEncode("State")

	corrDf := df.Correlation()
	corrDf.GetCol(1)
	// corrDf.Show(true)

	// df.Head(5)

}

func convertRound(x float64) int {
	if x < 0.5 {
		return 0
	}

	return 1
}
