package main

import (
	"ai/ai"
	"fmt"

	"github.com/kevinmtanadi/godf"
)

func main() {
	df := godf.ReadCSV("data/winequalityN.csv")
	// df.Scramble(5019320)
	// df.Head(30)

	shouldRun := true

	df.Info()
	df.OneHotEncode("type")
	// df.Scramble()

	df.Head()
	// df.Corr(

	df.GetRow(16, 17, 18).Show()

	if shouldRun {
		yHeaderName := "quality"

		y := df.GetCol(yHeaderName).ExtractData().([]interface{})
		df.DropCol(yHeaderName)
		x := df.Transpose().ExtractData().([][]interface{})

		castedX := Convert2DFloat(x)
		castedY := Convert1DFloat(y)

		trainSize := 0.8
		n := int(trainSize * float64(len(castedY)))
		xTrain := castedX[:n]
		yTrain := castedY[:n]
		xTest := castedX[n:]
		yTest := castedY[n:]

		fmt.Println(len(xTrain), len(xTest))
		fmt.Println(len(xTrain[0]), len(xTest[0]))

		bestAcc := 0.
		k := 0

		fmt.Println("Start predicting")
		for i := 1; i <= 5; i++ {
			model := ai.NewKNN(i, 2, xTrain, yTrain)

			yPred := model.Predict(xTest)
			correct := 0.
			for i, v := range yPred {
				if v == yTest[i] {
					correct++
				}
			}
			acc := correct / float64(len(yPred)) * 100
			fmt.Println("Accuracy on k: ", i, ": ", acc, "%")
			if acc > bestAcc {
				bestAcc = acc
				k = i
			}
		}

		fmt.Println("Best accuracy: ", bestAcc, " on k: ", k)

		bestModel := ai.NewKNN(1, 2, xTrain, yTrain)
		yPred := bestModel.Predict(xTest)

		comDf := godf.DataFrame(map[string]interface{}{
			"y_test": yTest,
			"y_pred": yPred,
		})

		comDf.Show()
	}

}
