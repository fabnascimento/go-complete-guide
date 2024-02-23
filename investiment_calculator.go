package main

import (
	"fmt"
	"math"
)

func main() {
	// explicitly setting type
	var investimentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investimentAmount * math.Pow((1+expectedReturnRate/100), years)

	fmt.Println("future value", futureValue)
}
