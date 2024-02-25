package cmdmanager

import "fmt"

type CMDManager struct {
}

// example to see why we better have structs for filemenager for example
func (cmd CMDManager) ReadLines() ([]string, error) {

	fmt.Println("Please enter your prices. Confirm every price with Enter and 0 for exit")

	var prices []string
	for {
		var price string
		fmt.Print("Price:")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)

	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
