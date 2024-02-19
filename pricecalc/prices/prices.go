package prices

import "fmt"

type TextIncludedPriceJob struct {
	TaxRate           float64
	InpputPrices      []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxrate float64) *TextIncludedPriceJob {

	return &TextIncludedPriceJob{
		InpputPrices: []float64{10.0, 20.0, 30.0},
		TaxRate:      taxrate,
	}
}

func (t TextIncludedPriceJob) Process() {

	result := make(map[string]float64)

	for _, p := range t.InpputPrices {

		result[fmt.Sprintf("%.2f", p)] = p * (1 + t.TaxRate)

	}
	fmt.Println(result)

}
