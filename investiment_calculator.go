package main

import (
	"fmt"
	"math"
)

func main() {
  const inflationRate = 2.5
	investimentAmount := 1000.0 // declares and defines a variable
	expectedReturnRate := 5.5
	var years float64 = 10

  fmt.Print("Investiment amount: ")
  fmt.Scan(&investimentAmount)

  fmt.Print("Expected return rate: ")
  fmt.Scan(&expectedReturnRate)

  fmt.Print("For how many years? ")
  fmt.Scan(&years)


	projectedFutureValue := investimentAmount * math.Pow((1+expectedReturnRate/100), years)
  futureRealValue := projectedFutureValue / math.Pow(1+inflationRate/100, years)
  
	fmt.Println("projected future value", projectedFutureValue)
	fmt.Println("future value", futureRealValue)
}
