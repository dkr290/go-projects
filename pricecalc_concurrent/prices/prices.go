package prices

import (
	"fmt"
	"pricecalc/conversion"
	"pricecalc/iomanager"
)

type TextIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InpputPrices      []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (t *TextIncludedPriceJob) LoadData() error {

	lines, err := t.IOManager.ReadLines()
	if err != nil {
		return err

	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	t.InpputPrices = prices

	return nil

}

func NewTaxIncludedPriceJob(iomn iomanager.IOManager, taxrate float64) *TextIncludedPriceJob {

	return &TextIncludedPriceJob{
		TaxRate:   taxrate,
		IOManager: iomn,
	}
}

func (t *TextIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {

	err := t.LoadData()

	if err != nil {

		errorChan <- err
		return
	}

	result := make(map[string]string)
	for _, p := range t.InpputPrices {

		taxIncludedPrice := p * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", p)] = fmt.Sprintf("%.2f", taxIncludedPrice)

	}
	t.TaxIncludedPrices = result
	t.IOManager.WriteResult(t)
	doneChan <- true

}
