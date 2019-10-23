package main

import (
	"cryptopals/supportfunctions"
	"fmt"
	"log"
)



func main() {

	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	byteInput, err := supportfunctions.HexStringToBytes(input)
	if err!=nil {
		log.Println("crashed due to hex input not hexy enough")
	}

	texts := make([]string, 128)
	decodedString := make([]string, 128)
	rowLength := len(byteInput)/len(string("A"))
	for letter := 0; letter < 128 ; letter++ {
		texts[letter]=string(supportfunctions.GetByteListFromClonedString(rowLength, string(letter)))
		msg := supportfunctions.XOROnBytes([]byte(byteInput), []byte(texts[letter]))
		decodedString[letter] = string(msg)
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
