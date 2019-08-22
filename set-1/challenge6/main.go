package main

import (
	"errors"
	// "../challenge2"
	// "encoding/hex"
	"fmt"
	// "strings"
	"math/bits"
)

func main () {
	fmt.Println(CumputeHammingWithChallege2("this is a test","wokka wokka!!!"))
	
}











func ComputeHammingDistance(str1 string, str2 string) (result int) {
	bytes1 := []byte(str1)
	bytes2 := []byte(str2)

	for index, bits1 := range bytes1 {
		bits2 := bytes2[index]
		result += bits.OnesCount(uint(bits1) ^ uint(bits2))
	}
	return result
}
func CumputeDistance(decoded_hex_string string, decoded_hex_string2 string) (result int, err error) {
	if len(decoded_hex_string) != len(decoded_hex_string2) {
		err = errors.New("Hex strings are not of equal lenght")	
		return	
	}

	for i:=0; i<len(decoded_hex_string); i++ {
		result+= bits.OnesCount(uint(decoded_hex_string[i])^uint(decoded_hex_string2[i]))
	}
	return
}