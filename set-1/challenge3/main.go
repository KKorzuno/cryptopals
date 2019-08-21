package main

import (
	"../challenge2"
	"encoding/hex"
	"fmt"
	"strings"
)

func englishCount(input string) (count int) {
	input = strings.ToLower(input)
	for _ , element := range "etaoin shrdlu" {
		count += strings.Count(input, string(element))
	}
	//fmt.Println(input)
	return
}


func main() {

	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	//fmt.Println(len(input))
	texts := make([]string, 'z'-'a'+1)
	decodedString := make([]string, 'z'-'a'+1)
	//fmt.Println('z' - 'a' + 1)
	//fmt.Println('a')
	for letter := 0; letter < 'Z'-'A'+1; letter++ {

		for i := 0; i < len(input)/len(hex.EncodeToString([]byte(string(letter+'A')))); i++ {
			//fmt.Printf("\nletter: %v, i:%v\n", letter, i)
			texts[letter] += hex.EncodeToString([]byte(string(letter + 'A')))

		}
		//fmt.Println(len([]byte(texts[letter])))
		//fmt.Printf("contents of byte: %v\n",[]byte(texts[letter]))
		//fmt.Printf("contents of hex(byte):%s\n",texts[letter])
		msg, err := challenge2.Fixed_XOR_on_hex_strings(input, texts[letter])
		if err != nil {
			fmt.Println("PROBLEMS with XOR")
			return
		}
		str2, err2 := hex.DecodeString(msg)
		if err2 != nil {
			fmt.Println("PROBLEMS with decoding the outcome hex")
			return
		}
		decodedString[letter] = string(str2)
		//fmt.Println(decodedString[letter])
	}

	var maxCount int
	var bestString string
	for _ , element := range decodedString {
		currentCount := englishCount(element)
		if currentCount > maxCount {
			bestString = element
			maxCount = currentCount
		}
		fmt.Println(currentCount)
		fmt.Println(bestString)
	}
}
