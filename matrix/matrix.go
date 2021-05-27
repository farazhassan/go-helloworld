package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	Data [][]string
}

func (m Matrix) Echo() string {
	var result string
	for _, row := range m.Data {
		result = fmt.Sprintf("%s%s\n", result, strings.Join(row, ","))
	}
	return result
}

func (m Matrix) Invert() string {
	inverted := transpose(m.Data)
	var result string
	for _, row := range inverted {
		result = fmt.Sprintf("%s%s\n", result, strings.Join(row, ","))
	}
	return result
}

func transpose(data [][]string) [][]string {
	xl := len(data[0])
	yl := len(data)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = data[j][i]
		}
	}
	return result
}

func (m Matrix) Flatten() string {
	var result []string
	for _, row := range m.Data {
		for _, item := range row {
			result = append(result, item)
		}
	}
	return strings.Join(result, ",")
}

func (m Matrix) Sum() (int, error) {
	result := 0
	for _, row := range m.Data {
		for _, val := range row {
			number, err := strconv.Atoi(val)
			if err != nil {
				return -1, errors.New("Non number value found.")
			}
			result += number
		}
	}
	return result, nil
}

func (m Matrix) Multiply() (int, error) {
	result := 1
	for _, row := range m.Data {
		for _, val := range row {
			number, err := strconv.Atoi(val)
			if err != nil {
				return -1, errors.New("Non number value found.")
			}
			result *= number
		}
	}
	return result, nil
}
