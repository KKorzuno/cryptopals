package main

import (
	"cryptopals/supportfunctions"
	"fmt"
	"log"
)



func main() {

	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	byteInput, err := supportfunctions.HexStringToBytes(input)
	if err!=nil {log.Println("crashed due to hex input not hexy enough")}
	//fmt.Println(len(input))
	texts := make([]string, 128)
	decodedString := make([]string, 128)
	//fmt.Println('z' - 'a' + 1)
	//fmt.Println('a')
	for letter := 0; letter < 128 ; letter++ {

		for i := 0; i < len(byteInput)/len(string(letter+'A')); i++ {
			//fmt.Printf("\nletter: %v, i:%v\n", letter, i)
			texts[letter] += string(letter + 'A')

		}
		//fmt.Println(len([]byte(texts[letter])))
		//fmt.Printf("contents of byte: %v\n",[]byte(texts[letter]))
		//fmt.Printf("contents of hex(byte):%s\n",texts[letter])
		msg := supportfunctions.XOROnBytes([]byte(byteInput), []byte(texts[letter]))

		decodedString[letter] = string(msg)
		//fmt.Println(decodedString[letter])
	}

	var maxCount int
	var bestString string
	for _ , element := range decodedString {
		currentCount := supportfunctions.EnglishCount(element)
		if currentCount > maxCount {
			bestString = element
			maxCount = currentCount
		}
		//fmt.Println(currentCount)
	
	}
	fmt.Println(bestString)
}
