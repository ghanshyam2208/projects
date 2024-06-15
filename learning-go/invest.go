package main

import (
	"fmt"
)

func main() {
	monthlyInvestment := 25000.0
	monthlyInterestRate := 5.0
	initialInvestment := 1000000.0

	totalMonths := 60

	currentInvestmentValue := initialInvestment
	totalInvestment := initialInvestment

	for month := 1; month <= totalMonths; month++ {
		// Calculate the monthly profit or loss based on the current investment value
		monthlyProfitLoss := (monthlyInterestRate * currentInvestmentValue) / 100
		currentInvestmentValue += monthlyProfitLoss

		// Calculate the growth
		growth := currentInvestmentValue - totalInvestment

		// Print the details for the current month
		fmt.Printf("Current month %d | Total invested %.2f | Current value %.2f | Growth %.2f \n", month, totalInvestment, currentInvestmentValue, growth)

		// Add the monthly investment to the current investment value and total investment
		currentInvestmentValue += monthlyInvestment
		totalInvestment += monthlyInvestment
	}
}
