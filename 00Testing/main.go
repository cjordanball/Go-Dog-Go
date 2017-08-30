package main

import (
	"fmt"
	"time"
)

func main() {
	rightNow := time.Now()
	fmt.Println(rightNow.Month())
}
