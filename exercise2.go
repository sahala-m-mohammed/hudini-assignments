// Objective:Create a command-line interface (CLI) application in Go that allows users to input a block of text and
// then calculates the frequency of each word in the text. This project will help you understand and
// implement basic Go concepts like variables, loops, conditionals, maps, functions, and strings.
// Requirements:
// Input Text: Allow users to input a block of text.
// Process Text: Count the frequency of each word in the text.
// Display Frequencies: Display the word frequencies in a readable format.
// Exit: Allow the user to exit the application.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	m := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	v := strings.TrimSpace(input)
	slice := strings.Split(v, " ")
	for i := 0; i < len(slice); i++ {
		_, exists := m[slice[i]]
		if exists { 
			m[slice[i]]++

		} else {
			m[slice[i]] = 1
		}
		
		}
	fmt.Println(m)
	}