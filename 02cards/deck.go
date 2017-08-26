package main

import "fmt"
// create a new type of "deck"
// which is a slice of strings
type deck []string 

func (dude deck) print() {
	for ind, val := range dude {
		fmt.Println(ind, val)
	}
}