//expected output 
// 10 --> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
//  1 --> [1]
 
package main
import "fmt"

func main(){
	fmt.Println(monkeyCount(10))
}
func monkeyCount(n int) []int {
   // Your code here, happy coding!
  var arr=[] int{}
  for i:=0;i<n;i++{
    arr=append(arr,i+1)
  }
  return arr
}

// MakeNegative(1)    // return -1
// MakeNegative(-5)   // return -5
// MakeNegative(0)    // return 0

func MakeNegative(x int) int {
  if x>0{
    return x*-1
  }else if x<0{
    return x
  }else{
    return 0
  }
}

// In this simple exercise, you will build a program that takes a value, integer , and returns a list of its multiples up to another value, limit . If limit is a multiple of integer, it should be included as well. There will only ever be positive integers passed into the function, not consisting of 0. The limit will always be higher than the base.

// For example, if the parameters passed are (2, 6), the function should return [2, 4, 6] as 2, 4, and 6 are the multiples of 2 up to 6.


package main

func FindMultiples(integer, limit int) []int {
  // Your code here!
  var arr=[]int{}
  n:=limit/integer
  for i:=1;i<=n;i++{
    arr=append(arr,i*integer)
  }
  return arr
}

// countBy(1,10)  should return  {1,2,3,4,5,6,7,8,9,10}
// countBy(2,5)  should return {2,4,6,8,10}

package main


func CountBy(x, n int) []int {
  var arr=[]int{}
  for i:=1;i<=n;i++{
    arr=append(arr,i*x)
  }
  return arr
}

// n = 0  ==> [1]        # [2^0]
// n = 1  ==> [1, 2]     # [2^0, 2^1]
// n = 2  ==> [1, 2, 4]  # [2^0, 2^1, 2^2]

package main

func PowersOfTwo(n int) []uint64 {
  // your code goes here
  var arr=[]uint64{1}
  var val uint64=1
  for i:=1;i<=n;i++{
    val=val*2
    arr=append(arr,val)
  }
  return arr
}


// Our football team has finished the championship.

// Our team's match results are recorded in a collection of strings. Each match is represented by a string in the format "x:y", where x is our team's score and y is our opponents score.

// For example: ["3:1", "2:2", "0:1", ...]

// Points are awarded for each match as follows:

// if x > y: 3 points (win)
// if x < y: 0 points (loss)
// if x = y: 1 point (tie)
// We need to write a function that takes this collection and returns the number of points our team (x) got in the championship by the rules given above.

// Notes:

// our team always plays 10 matches in the championship
// 0 <= x <= 4
// 0 <= y <= 4

package kata
import "strings"

func Points(games []string) int {
  // your code here!
  var x int
  for i:=0;i<len(games);i++{    
    val:=strings.Split(games[i],":")
    if(val[0]>val[1]){
      x=x+3
    }else if val[0]==val[1]{
      x=x+1
    }else{
      
    }
  }
  return x
}

// Given two integers a and b, which can be positive or negative, find the sum of all the integers between and including them and return it. If the two numbers are equal return a or b.

// Note: a and b are not ordered!

// Examples (a, b) --> output (explanation)
// (1, 0) --> 1 (1 + 0 = 1)
// (1, 2) --> 3 (1 + 2 = 3)
// (0, 1) --> 1 (0 + 1 = 1)
// (1, 1) --> 1 (1 since both are same)
// (-1, 0) --> -1 (-1 + 0 = -1)
// (-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)
// Your function should only return a number, not the explanation about how you get that number.

package kata


func GetSum(a, b int) int {
  var sum int
  if a==b{
    sum=a
  }else if a<b{
    for i:=a;i<=b;i++{
      sum=sum+i
    }
  }else{
      for i:=b;i<=a;i++{
        sum=sum+i
      }
    }
  return sum
}


// Simple, given a string of words, return the length of the shortest word(s).

// String will never be empty and you do not need to account for different data types.

package kata
import "strings"

func FindShort(s string) int {
  // your code
  var shortest int
  val:=strings.Split(s," ")
  shortest=1000
  
  for i:=0;i<len(val);i++{
    if (len(val[i])<shortest){
      shortest=len(val[i])
    }
  }
  return shortest
}

//Complete the square sum function so that it squares each number passed into it and then sums the results together.


package kata

func SquareSum(numbers []int) int {
    // your code here
  var sum int
  for i:=0;i<len(numbers);i++{
    square:=numbers[i]*numbers[i]
    sum=sum+square
  }
  return sum
}


// 7  -> 3 (because odd numbers below 7 are [1, 3, 5])
// 15 -> 7 (because odd numbers below 15 are [1, 3, 5, 7, 9, 11, 13])

package kata

func OddCount(n int) int{
  //your code here
  return n/2
}

