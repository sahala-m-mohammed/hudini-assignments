package main

import "fmt"

func main() {
    //1
    // str := []int{10, 20, 30, 40, 50}
    // fmt.Printf("CalculateAverage({10, 20, 30, 40, 50})= %d \n", CalculateAverage(str))
    // //2
    // fmt.Printf("Checkage(25)= %s\n", checkAge(25))
    // //3
    // strl := "hello"
    // fmt.Printf("reverseString(hello)= %s\n", reverseString(strl))
    // //4
    // fmt.Printf("largest({10, 20, 30, 40, 50})= %d", findLargestNumber(str))
    var counter =Counter{countn:0}
   
    counter=counter.add();
    counter.add();
    counter=counter.display();  // Output: 2
    counter=counter.subtract();
    counter.display();  // Output: 1
    counter.reset();
    counter.display();
}
 
type Counter struct{
    countn int
    }
func Counterfn(count Counter) Counter{
    func add(count Counter) counter{
        return count.countn++
    }
    func subtract(count Counter) counter{
        return count.countn--
    }
    func (count Counter) display() Counter {
        fmt.println(count.countn)
    }
    func (count Counter) reset() Counter {
        count.countn=0
        return count.countn
    }      
}
has context menu
	func display(counter Counter) Counter{
		return counter.counter
	}
	func reset(counter Counter) Counter{
		counter.counter=0
	
}




