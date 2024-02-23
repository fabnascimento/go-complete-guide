package main

import (
	"fmt"
	"math"
)

func main() {
	var investimentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investimentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))

	fmt.Println("future value", futureValue)
}
