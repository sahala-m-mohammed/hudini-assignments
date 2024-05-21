package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}





LASAGNA functions prob

n =len(numberOfLayers)
    for i:=0;i<n;i++{
        if avgPrepTime==0{
        	time= n* 2
       	}else{
        	time= n*avgPrepTime
    }