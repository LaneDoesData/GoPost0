package Utils

import "strconv"

func String2Int(s string) int {
	intVar, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return intVar
}

func String2float(s string) float64 {
	floatVar, err := strconv.ParseFloat(s,64)
	if err != nil {
		return 0.0
	}
	return floatVar
}