// 1. Calculate Average
// Objective: Write a function that takes an array of numbers and returns the average. Use loops and basic arithmetic.
// Function signature:
package main

import "fmt"

func main() {
	var numbers = [] int {10, 20, 30, 40, 50}
	fmt.Println(calculateAverage(numbers))
}

func calculateAverage(a []int) float64 {
	var sum int
	var average float64
    for i:=0;i<len(a);i++{
        sum= sum + a[i]
    }
    average = float64(sum)/float64(len(a))
    return average
}
    // Write your code here to calculate and return the average of the array elements.

 
// Example usage:
 // Expected output: 30
 
// 2. Check Age
// Write a function that takes an age and prints whether the person is a minor, a young adult, or an adult.
// Use conditional statements.
// Function signature:
package main

import "fmt"

func main() {
	var age int=25
	fmt.Println(checkAge(age))
}

func checkAge(age int) string {
	if age<18 && age>=0{
		return "minor"
	}else if age>=18 && age <=35{
		return "young adult"
	}else{
		return "adult"
	}
}

    // Write your code here to determine and print if the age corresponds to a minor, a young adult, or an adult.

 
// Example usage:
// Possible output: young adult
 
// 3. Reverse String
// Objective: Create a function that reverses a string. This will demonstrate basic string manipulation and for loops.
// Function signature:
 // Write your code here to reverse and return the string.
import "fmt"

func main() {
	var str string = "hello"
	fmt.Println(reverseString(str))
}

func reverseString(stri string) string {
	var newString string
	for i:=len(stri)-1;i>=0;i--{
		newString+=string(stri[i])
	}
	return newString
}
   

 
// Example usage:
// Expected output: "olleh"
 
// 4. Find Largest Number
// Objective: Write a function that takes an array of numbers and returns the largest number.
// Use loops and conditional statements to solve the problem.
// Function signature:
// Write your code here to find and return the largest number in the array.
package main

import "fmt"

func main() {
	var numbers = [] int {10, 20, 30, 40, 50}
	fmt.Println(findLargestNumber(numbers))
}

func findLargestNumber(large []int) int {
	var largest int = 0
    for i:=0;i<len(large);i++{
        if largest<large[i]{
			largest= large[i]
		}
	}
    return largest
}
 
// Example usage:
//findLargestNumber([10, 20, 30, 40, 50]); // Expected output: 50
 
// 5. Simple Counter
// Create an object that acts as a counter with methods to add, subtract, and reset the count.
// Demonstrate object methods and the use of this.
// Object definition:
const counter = {
 
};
 
// Example usage:
counter.add();
counter.add();
counter.display();  // Output: 2
counter.subtract();
counter.display();  // Output: 1
counter.reset();
counter.display();  // Output: 0
 