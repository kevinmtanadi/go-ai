package dataframe

import (
	"ai/formula"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type DataFrame struct {
	Header []string
	Data   [][]interface{}
}

func (d *DataFrame) Show(headerAsRow bool) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)
	for idx, row := range d.Data {
		if idx == 0 && d.Header != nil {
			fmt.Fprintln(w, fmt.Sprintf("\t%s", strings.Join(d.Header, "\t")))
		}
		var line string
		if headerAsRow {
			// fmt.Println(len(d.Data))
			// fmt.Println(len(d.Data[0]))
			// fmt.Println(d.Data[0])
			// if len(d.Data) != len(d.Data[0]) {
			// 	panic("Dimension mismatch")
			// }
			line = fmt.Sprintf("%s\t%s", d.Header[idx], strings.Join(stringify(row), "\t"))
		} else {
			line = fmt.Sprintf("%d\t%s", idx+1, strings.Join(stringify(row), "\t"))
		}
		fmt.Fprintln(w, line)
	}
	w.Flush()
}

func (d *DataFrame) Head(n int) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)
	for idx, row := range d.Data {
		if idx == 0 && d.Header != nil {
			fmt.Fprintln(w, fmt.Sprintf("\t%s", strings.Join(d.Header, "\t")))
		}
		line := fmt.Sprintf("%d\t%s", idx+1, strings.Join(stringify(row), "\t"))
		fmt.Fprintln(w, line)
		if idx == n-1 {
			break
		}
	}
	w.Flush()
}

func stringify(data []interface{}) []string {
	line := make([]string, len(data))

	for idx, x := range data {
		if i, ok := x.(int); ok {
			line[idx] = strconv.Itoa(i)
		} else if f, ok := x.(float64); ok {
			line[idx] = strconv.FormatFloat(f, 'f', -1, 64)
		}
	}

	return line
}

func (d *DataFrame) ReadCSV(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	line := 0
	for {
		data, err := reader.Read()
		if err != nil {
			break
		}

		if line == 0 {
			d.Header = data
		} else {
			castedDatas := make([]interface{}, len(data))
			for i, s := range data {
				castedData := castDataType(s)
				castedDatas[i] = castedData
			}
			d.Data = append(d.Data, castedDatas)
		}
		line++
	}
}

func (d *DataFrame) InsertFloat(data [][]float64) {
	length := len(data[0])

	for _, line := range data {
		interfaceSlice := make([]interface{}, length)
		for i, d := range line {
			interfaceSlice[i] = d
		}
		d.Data = append(d.Data, interfaceSlice)
	}

}

func castDataType(s string) interface{} {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	return s
}

// GetRow
//
// Returns a row of given row number
func (d *DataFrame) GetRow(rowNum int) []interface{} {
	return d.Data[rowNum]
}

// GetCol
//
// Returns a column of given column number
func (d *DataFrame) GetCol(colNum int) []interface{} {
	if len(d.Data) == 0 {
		panic("No data in dataframe")
	}
	data := make([]interface{}, len(d.Data))

	for i := range d.Data {
		data[i] = d.Data[i][colNum]
	}

	return data
}

// OneHotEncode
//
// Encode string data into int
// It takes the header's name as an argument
func (d *DataFrame) OneHotEncode(headerName string) {
	encodeMap := make(map[string]float64)

	var colNum int
	for i, h := range d.Header {
		if h == headerName {
			colNum = i
		}
	}

	for i, v := range d.GetCol(colNum) {
		if _, ok := encodeMap[v.(string)]; !ok {
			encodeMap[v.(string)] = float64(len(encodeMap))
		}
		d.Data[i][colNum] = encodeMap[v.(string)]
	}

}

func (d *DataFrame) Correlation() DataFrame {
	corrTable := [][]float64{}

	length := len(d.Data[0])

	for i := 0; i < length; i++ {
		corrLine := make([]float64, length)
		for j := 0; j < length; j++ {
			corrLine[j] = formula.Correlation(castFloat(d.GetCol(i)), castFloat(d.GetCol(j)))
		}
		corrTable = append(corrTable, corrLine)
	}

	corrDf := DataFrame{}
	corrDf.InsertFloat(corrTable)
	corrDf.setHeader(d.Header)
	corrDf.Show(true)

	return corrDf
}

func (d *DataFrame) setHeader(headers []string) {
	for _, h := range headers {
		d.Header = append(d.Header, h)
	}
}

func castFloat(x []interface{}) []float64 {
	data := make([]float64, len(x))
	for i, v := range x {
		if f, ok := v.(float64); ok {
			data[i] = f
		}
	}

	return data
}
