package main
import "fmt"

//Slices in go

func main() {
	var deck = []string {"Card 1"}
	deck = append(deck, "Card 2")
	fmt.Println(deck)

	var cards [] string;
	cards = append(cards, "Card 3")
	cards = append(cards, "Card 4")
	fmt.Println(cards)

	phones:= []string{"iPhone", "Samsung", "OnePlus"}
	phones = append(phones, "Google Pixel")
	fmt.Println(phones)

	for i, phone := range phones {
		fmt.Println(i, phone)
	}
}
