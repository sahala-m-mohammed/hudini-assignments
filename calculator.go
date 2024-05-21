package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {
	//input: 2 values of type float64
	//give switch for each operand case
	//return : value after operation
	var result float64
	switch operand {
	case "+":
		result = float64(value1 + value2)
	case "-":
		result = float64(value1 - value2)
	case "*":
		result = float64(value1 * value2)
	case "/":
		result = float64(value1 / value2)
	default:
		fmt.Println("Invalid operand", nil)
	}
	return result, nil
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
	if len(parts) == 3 {
		// TODO: Take all 3 values and pass to calculator function
		a, _ := strconv.ParseFloat(parts[0], 64)
		b, _ := strconv.ParseFloat(parts[2], 64)
		result, _ := Calculator(a, b, parts[1])
		// TODO: Print results
		fmt.Printf("Result: %v\n", result)
	} else {
		fmt.Println("Invalid format")
	}
}
