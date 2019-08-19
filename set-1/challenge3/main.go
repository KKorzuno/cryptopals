package main

import (
	"../challenge2"
	"encoding/hex"
	"fmt"
)

func main() {

	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	fmt.Println(len(input))
	texts := make([]string, 'z'-'a'+1)
	
	/*
	fmt.Println('z' - 'a' + 1)
	fmt.Println('a')
	for letter := 0; letter < 'z'-'a'+1; letter++ {

		for i := 0; i < len(input); i++ {
			fmt.Printf("\nletter: %v, i:%v\n", letter, i)
			texts[letter] += string(letter + 'a')

		}
		fmt.Println(hex.EncodeToString([]byte(texts[letter])))
		fmt.Println("ABC")
		fmt.Printf("lenght of input: %v length of hex_text: %v", len(input), len(hex.EncodeToString([]byte(texts[letter]))))
		msg, err := challenge2.Fixed_XOR_on_hex_strings(input, hex.EncodeToString([]byte(texts[letter])))
		if err != nil {
			fmt.Println("PROBLEMS with XOR")
			return
		}
		str2, err2 := hex.DecodeString(msg)
		if err2 != nil {
			fmt.Println("PROBLEMS with decoding the outcome hex")
			return
		}
		fmt.Println(string(str2))
	}
	*/
}
