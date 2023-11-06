package main

import (
	"ai/ai"
	"ai/dataframe"
	"ai/dataset/preprocessing"
	"fmt"
)

func main() {
	df := dataframe.DataFrame{}
	df.ReadCSV("data/diabetes.csv")

	// corrDf := df.Correlation()
	// corrDf.Show(true)

	lowCorr := []interface{}{
		"BloodPressure",
		"SkinThickness",
	}

	df.Drop(lowCorr...)

	y := df.GetCol("Outcome").GetFloatData()[0]

	df.Drop("Outcome")

	df.Head(1)

	x := df.GetFloatData()

	trainData := int(0.8 * float64(len(x)))

	X_train := x[:trainData]
	y_train := y[:trainData]
	x_test := x[trainData:]
	y_test := y[trainData:]

	// X_train := [][]float64{
	// 	{0.08, 0.72}, {0.26, 0.58},
	// 	{0.45, 0.15}, {0.60, 0.30},
	// 	{0.10, 1.0}, {0.35, 0.95},
	// 	{0.70, 0.65}, {0.92, 0.45},
	// }

	// y_train := []float64{
	// 	1, 1, 1, 1, 0, 0, 0, 0,
	// }

	// x_test := [][]float64{
	// 	{0.1, 0.9},
	// 	{0.4, 0.2},
	// 	{0.8, 0.5},
	// 	{0.2, 0.6},
	// }

	// y_test := []float64{0, 1, 0, 1}

	standardizer := preprocessing.Standardizer{}
	X_train = standardizer.FitTransform(X_train)
	x_test = standardizer.FitMultiple(x_test)

	fmt.Println(X_train[0])

	// model := ai.LogisticRegression{
	// 	Epochs:       10000,
	// 	LearningRate: 0.075,
	// }

	// model.Train(X_train, y_train)
	// fmt.Println(model.W0)
	// fmt.Println(model.Wi)

	// y_pred := model.Predict(x_test)

	accuracies := []float64{}
	highestAccuracy := 0.
	bestK := 0

	for k := 1; k < 300; k++ {

		model := ai.NewKNN(k, 2, X_train, y_train)

		y_pred := model.Predict(x_test)

		result := dataframe.DataFrame{}

		result.Header = []string{
			"Prediction",
			"Actual",
		}

		for i := range y_pred {
			result.Data = append(result.Data, []interface{}{
				convertRound(y_pred[i]),
				y_test[i],
			})
		}

		// result.Show()

		correct := 0.

		for _, row := range result.Data {
			if row[0] == row[1] {
				correct++
			}
		}

		accuracy := correct / float64(len(result.Data))
		accuracies = append(accuracies, accuracy)

		if accuracy > highestAccuracy {
			highestAccuracy = accuracy
			bestK = k
		}
		// fmt.Println(correct / float64(len(result.Data)))
	}

	fmt.Println(bestK, highestAccuracy)

	// df.Head(5)
}

func convertRound(x float64) float64 {
	if x < 0.5 {
		return 0
	}

	return 1
}
