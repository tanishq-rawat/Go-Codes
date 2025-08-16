package main

import "fmt"

func getCard() string {
	return sampleFunction()
}

func main() {
	var card string = getCard()
	fmt.Println("My card is:", card)
}