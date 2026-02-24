package main

import (
	"fmt"
)

func main() {
	var investmentAmt float64
	var salePrice float64
	var amt float64
	fmt.Print("\nEnter the Actual Product Cost = ")
	fmt.Scanln(&investmentAmt)

	fmt.Print("\nEnter the Sale Price = ")
	fmt.Scanln(&salePrice)

	if salePrice > investmentAmt {
		amt = salePrice - investmentAmt
		fmt.Println("Total Profit = ", amt)
	} else if investmentAmt > salePrice {
		amt = investmentAmt - salePrice
		fmt.Println("Total Loss = ", amt)
	} else {
		fmt.Println("No Profit No Loss")
	}

}
