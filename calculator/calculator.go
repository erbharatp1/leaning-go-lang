package calculator

import (
	"damo-go/greetings"
	"fmt"
)

func main() {
	var productAmt float64
	var productSalePrice float64
	var amt float64
	fmt.Println(greetings.Hello("Bharat"))
	fmt.Print("\nEnter the Actual Product Cost = ")
	fmt.Scanln(&productAmt)

	fmt.Print("\nEnter the Sale Price = ")
	fmt.Scanln(&productSalePrice)

	if productSalePrice > productAmt {
		amt = productSalePrice - productAmt
		fmt.Println("Total Profit = ", amt)
	} else if productAmt > productSalePrice {
		amt = productAmt - productSalePrice
		fmt.Println("Total Loss = ", amt)
	} else {
		fmt.Println("No Profit No Loss")
	}

	greetings.Hello("Bharat")
}
