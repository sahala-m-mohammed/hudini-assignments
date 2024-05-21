package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {
	// TODO: Implement Calculator function
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 10 + 20): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	// TODO: Check if exact 3 parts are given. If not, throw an error

	// TODO: Take all 3 values and pass to calculator function

	// TODO: Print results
	fmt.Printf("Result: %v\n", result)
}
