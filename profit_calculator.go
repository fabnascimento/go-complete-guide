// Build a Profit Calculator

// ask for revenue, expenses & tax rate
// calculate earnings before TAX (EBT) and
// earnings after tax (Profit)

// calculate ratio (EBT / Profit)

// output EBT, profite and the ration
package main

import "fmt"

func getUserInput(prompt string) float64 {
  var value float64

  fmt.Print(prompt)
  fmt.Scan(&value)

  return value
}

func calculateEBT(revenue, expenses float64) float64 {
  return revenue - expenses
}

func calculateProfit(ebt, taxRate float64) float64 {
  return ebt * (1 - taxRate/100) 
}

func calculateEBTProfitRatio(ebt, profit float64) float64 {
  return ebt/profit
}

func main() { 
  revenue := getUserInput("Enter your revenue: ")
  expenses := getUserInput("Enter expenses: ")
  taxRate := getUserInput("Enter the tax rate: ") 
  
  EBT := calculateEBT(revenue, expenses)
  profit := calculateProfit(EBT, taxRate)
  ratio := calculateEBTProfitRatio(EBT, profit)

  fmt.Println("EBT:", EBT)
  fmt.Println("Profit:", profit)
  fmt.Println("Ration:", ratio)
}

