package main

import (
	"fmt"
	"pricecalc/filemanager"
	"pricecalc/prices"
)

func main() {

	var taxRates = []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	errorChans := make([]chan error, len(taxRates))

	// we can swap with command manager
	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdmng := cmdmanager.New()
		//without interfaces we cannot replace fm
		//we need that both filemanager and command manager are accepted
		// with interfaces it can be passed both filemanager or iomanager
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[i], errorChans[i])

		// if err != nil {
		// 	fmt.Println("could not process job")
		// 	fmt.Println(err)
		// }

	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done")
		}

	}

}
