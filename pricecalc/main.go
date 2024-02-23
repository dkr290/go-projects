package main

import (
	"fmt"
	"pricecalc/filemanager"
	"pricecalc/prices"
)

func main() {

	var taxRates = []float64{0, 0.07, 0.1, 0.15}
	// we can swap with command manager
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices1.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdmng := cmdmanager.New()
		//without interfaces we cannot replace fm
		//we need that both filemanager and command manager are accepted
		// with interfaces it can be passed both filemanager or iomanager
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}

	}

}
