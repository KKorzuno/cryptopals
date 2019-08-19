package main

import (
	"../challenge2"
	"fmt"
	"encoding/hex"
)

func main() {

	for i := '0'; i <= '9'; i++ {
		str, err  := XORStringAgainstSingleCharacter("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",string(i))
		for i=0; i< len(str);i++ {
			fmt.Print("%s", hex.DecodeString(str[i]))
		}
		if err != nil {}
	}
	for i := 'a'; i <= 'f'; i++ {
		fmt.Println(XORStringAgainstSingleCharacter("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",string(i)))
	}


}

func XORStringAgainstSingleCharacter(str string, char string) (output string, err error) {
	// hexedLetter := hex.EncodeToString([]byte(char))

	var longListOfLetters string
	for i :=0 ;i < len(str); i++ {
		longListOfLetters+=(char)
	}

	// var longListOfHexLetters string
	// for i :=0 ;i < len(str); i++ {
	// 	longListOfHexLetters+=hexedLetter
	// }

	// longListOfLettersHex := hex.EncodeToString([]byte(longListOfLetters))
	fmt.Println(char)
	// fmt.Println(hexedLetter)
	// fmt.Println(longListOfHexLetters)
	// fmt.Println(longListOfLettersHex)
	fmt.Println(len(str))
	// fmt.Println(len(longListOfLettersHex))
	return 	challenge2.Fixed_XOR_on_hex_strings(str,longListOfLetters)
}
