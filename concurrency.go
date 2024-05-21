// Objective:
// Learn how to send and receive values using channels.

// Task:

// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// The main function should receive the message from the channel and print it.
// Hints:

// Use an unbuffered channel for simplicity.

//package main

// import "fmt"

// func main() {
// 	msg := make(chan string)
// 	go func() {
// 		msg <- "Hello World!"
// 	}()

// 	m:= <- msg
// 	fmt.Println(m)
// }

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to create and use goroutines.

// Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.

// package main

// import (
// 	"fmt"
// 	"time"
// 	"sync"
// )
// var wg sync.WaitGroup

// func add1(){
// 	for i:=0;i<5;i++{
// 		time.Sleep(time.Second)
// 		fmt.Println(i+1)
// 	}
// 	wg.Done()
// }
// func add2(){
// 	for i:=0;i<5;i++{
// 		time.Sleep(time.Second)
// 		fmt.Println(i+1)
// 	}
// 	wg.Done()
// }
// func add3(){
// 	for i:=0;i<5;i++{
// 		time.Sleep(time.Second)
// 		fmt.Println(i+1)
// 	}
// 	wg.Done()
// }

// func main() {
// 	wg.Add(3)
// 	go add1()
// 	go add2()
// 	go add3()

// 	wg.Wait()

//}
// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// func printNo(wg *waitGroup){
// 	for i:=0;i<5;i++{
// 		fmt.Print(i+1)
// 	}
// 		wg.Done()
// }

// type waitGroup struct {
// 	count int
// }

// func (wg *waitGroup) Add() {
// 	wg.count++
// 	fmt.Println(wg.count)
// }
// func (wg *waitGroup) Done() {
// 	wg.count--
// 	fmt.Println(wg.count)
// }
// func (wg *waitGroup) Wait() {
// 	for{
// 		if wg.count == 0 {
// 			os.Exit(0)
// 		}
// 	}

// }

// func main() {
// 	wg := waitGroup{
// 		count:0,
// 	}

// 	wg.Add()
// 	go printNo(&wg)
// 	wg.Add()
// 	time.Sleep(time.Second)
// 	wg.Done()
// 	wg.Wait()
// }

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.
// package main

// import (
// 	"fmt"
// 	"time"
// )


// func hello(msg chan string) {
// 	msg <- "Hello "
// }

// func world(msg chan string) {
// 	time.Sleep(time.Second)
// 	msg <- "World"
// }

// func main() {
// 	message := make(chan string)
// 	go hello(message)
// 	go world(message)
// 	fmt.Print(<-message)
// 	fmt.Print(<-message)

// }

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to use channels for communication between goroutines.

// Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.

package main

import "fmt"

func OneToTen() int {
	for i:=0;i<11;i++{
		return i+1
	}
}

func main(){
	channel := make(chan int)
	
}

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use a buffered channel.

// Task:

// Write a program that creates a buffered channel with a capacity of 3.
// The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// Then, receive and print the integers from the channel.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:

// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.

// -------------------------------------------------------------------------------------

// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// Task:

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.
