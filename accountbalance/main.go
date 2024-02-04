package main

import (
	"accountbalance/pkg/fileops"
	"fmt"
	"os"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

var accountBalance float64

func main() {

	var err error
	fileops := fileops.NewFileOps(accountBalanceFile, accountBalance)

	accountBalance, err = fileops.GetFloatFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------------------")
		os.Exit(1)

	}
	fmt.Println("Reach out to the phone:", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your Choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount, must be > 0")
				continue
			}
			accountBalance += depositAmount
			fileops.Value = accountBalance
			fmt.Println("Balance Updated! New amount:", accountBalance)
			fileops.WriteFloatToFile()
		case 3:
			fmt.Print("Amount to withdraw:")
			var wAmount float64
			fmt.Scan(&wAmount)
			if wAmount <= 0 {
				fmt.Println("Invalid amount, must be > 0")
				continue
			}
			if wAmount > accountBalance {
				fmt.Println("Invalid amount , you cannot withdraw more of the total amount")
				continue
			}
			accountBalance -= wAmount
			fileops.Value = accountBalance
			fmt.Println("Your new amount is:", accountBalance)
			fileops.WriteFloatToFile()

		default:
			fmt.Println("GoodBye")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}

}
