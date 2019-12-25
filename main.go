package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

//const length = 12

func main() {

	length := flag.Int("num", 8, "Text to parse.")

	flag.Parse()

	//Aa-Zz
	charMap := append(charSlice(65, 90), charSlice(97, 122)...)
	//0-9
	charMap = append(charMap, charSlice(48, 57)...)
	//special char
	charMap = append(charMap, charSlice(33, 47)...)
	//special char
	charMap = append(charMap, charSlice(91, 96)...)
	//special char
	charMap = append(charMap, charSlice(58, 64)...)

	for i := 1; i <= *length; i++ {
		fmt.Printf("%c", charMap[randNum(len(charMap))])
	}

	fmt.Println()

}

func randNum(length int) int64 {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
	if err != nil {
		fmt.Println(err)
	}
	return num.Int64()
}

func charSlice(low int, high int) []int {
	var r []int
	for i := low; i <= high; i++ {
		r = append(r, int(i))
	}
	return r
}
