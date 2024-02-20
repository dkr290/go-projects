package prices

import (
	"fmt"
	"pricecalc/conversion"
	"pricecalc/filemanager"
)

type TextIncludedPriceJob struct {
	TaxRate           float64
	InpputPrices      []float64
	TaxIncludedPrices map[string]float64
}

func (t *TextIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return

	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.InpputPrices = prices

}

func NewTaxIncludedPriceJob(taxrate float64) *TextIncludedPriceJob {

	return &TextIncludedPriceJob{

		TaxRate: taxrate,
	}
}

func (t *TextIncludedPriceJob) Process() {

	t.LoadData()
	result := make(map[string]string)
	for _, p := range t.InpputPrices {

		taxIncludedPrice := p * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", p)] = fmt.Sprintf("%.2f", taxIncludedPrice)

	}
	fmt.Println(result)

}
