package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(s []string) ([]float64, error) {
	lFloats := make([]float64, len(s))
	for i, l := range s {
		floatPrice, err := strconv.ParseFloat(l, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float" + err.Error())
		}
		//prices = append(prices, floatPrice)
		lFloats[i] = floatPrice

	}

	return lFloats, nil
}
