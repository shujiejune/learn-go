package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for index, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			fmt.Println("Converting string to float failed.")
			return nil, err
		}
		floats[index] = floatVal
	}
	return floats, nil
}
