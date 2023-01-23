/*
Specification:
array with integers is given. Each integer represents colour of a sock. 
Return amount of pairs which can be formed if colours have to be matching.
Length of array is also provided separately as an argument, which is between 1 and 100.
*/

package main

import "fmt"

func main() {
	sockPile := []int{6,1,3,5,5,3,13,2,4,1,1,4,6,1,2}
	len := len(sockPile) //Golang has functionality to do just fine without knowing len in advance
	fmt.Println(sockMerchant(sockPile, len))
}

func sockMerchant(pile []int, len int) int {
	pairN := 0
	clusteredPile := make(map[int]int)

	for _, sockColour := range pile{
		if _, present := clusteredPile[sockColour]; present{
			clusteredPile[sockColour]++
		} else {
			clusteredPile[sockColour] = 1
		}
	}

	for _, sameColour := range clusteredPile{
		pairN += sameColour / 2
	}

	return pairN
}