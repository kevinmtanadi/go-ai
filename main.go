package main

import (
	"ai/dataframe"
	"ai/dataset/preprocessing"
	"fmt"
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

	// corrDf := df.Correlation()
	// corrDf.Show(true)

	df.Drop("Administration")
	df.Head()

	y := df.GetCol("Profit")

	fmt.Println(y)

	df.Drop("Profit")

	x := df.GetFloatData()
	fmt.Println(x)

	X_train := x[:40]
	y_train := y[:40]
	x_test := x[40:]
	y_test := y[40:]

	standardizer := preprocessing.Standardizer{}
	X_train = standardizer.FitTransform(X_train)
	x_test = standardizer.FitMultiple(x_test)

	// df.Head(5)

}

func convertRound(x float64) int {
	if x < 0.5 {
		return 0
	}

	return 1
}
