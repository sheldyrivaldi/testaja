package main

import "fmt"

func main() {
	var isSellerInit bool
	buySell := ""
	SELLER := "S"
	BUYER := "B"

	if buySell == SELLER {
		isSellerInit = true
	} else if buySell == BUYER {
		isSellerInit = false
	}

	fmt.Println(isSellerInit)
}
