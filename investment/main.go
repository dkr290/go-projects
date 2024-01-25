package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5

	var investmentAmount, years float64
	expectationRate := 5.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("The rate amount in percent: ")
	fmt.Scan(&expectationRate)

	futureValue := investmentAmount * math.Pow(1+expectationRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
