package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var investmentAmt float64
	var expectedReturnRate float64
	var years float64
	fmt.Print("Enter investment Amt: ")
	fmt.Scan(&investmentAmt)
	fmt.Print("Enter investment Years: ")
	fmt.Scan(&years)
	fmt.Print("Enter investment expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)
	featureInvestment := investmentAmt * math.Pow(1+expectedReturnRate/100, years)
	log.Println("Investment calculation completed", featureInvestment)

}
