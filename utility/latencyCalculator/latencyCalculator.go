package latencycalculator

import (
	"fmt"
	"math/rand"
	"time"
)

// Calculate Generates a random number in range of two input numbers
func Calculate(num1 float64, num2 float64) int {
	rand.Seed(time.Now().UnixNano())
	min := int(num1 * 1000)
	max := int(num2 * 1000)
	random := rand.Intn(int((max - min + 1) + min))
	fmt.Println(random, num1, num2)
	return random
}
