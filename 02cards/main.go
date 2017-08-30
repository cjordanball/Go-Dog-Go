package main

import "fmt"

// func main() {
// 	cards :=  newDeck()

// 	hand, remainingDeck := deal(cards, 5)

// 	hand.print()

// 	remainingDeck.print()
// }
func main() {

	title, pages := getBookInfo()

	fmt.Println(title);
	fmt.Println(pages);
}

func getBookInfo() (string, pages) {
	return "War and Peace", 1000
}