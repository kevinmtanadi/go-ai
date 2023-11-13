package main

func Convert2DFloat(x [][]interface{}) [][]float64 {
	floatArr := [][]float64{}
	for _, i := range x {
		casted := []float64{}
		for _, j := range i {
			switch v := j.(type) {
			case float64:
				casted = append(casted, v)
			case int:
				casted = append(casted, float64(v))
			}
		}
		floatArr = append(floatArr, casted)
	}

	return floatArr
}

func Convert1DFloat(x []interface{}) []float64 {
	floatArr := []float64{}
	for _, i := range x {
		switch v := i.(type) {
		case float64:
			floatArr = append(floatArr, v)
		case int:
			floatArr = append(floatArr, float64(v))
		}
	}

	return floatArr
}
