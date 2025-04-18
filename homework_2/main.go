package main

import "fmt"

func main() {
	fmt.Println("hello from main")
	myFunc()
	myMessage()
}

func myFunc() {
	fmt.Println("hello from my func!!")
}

func myMessage() {
	fmt.Println("My message")
	// Comment
}
