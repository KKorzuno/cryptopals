package main

import (
	"../challenge2"
	"encoding/hex"
	"fmt"
	"strings"
	"log"
	"bufio"
	"os"
)



func main() {
	file, err := os.Open("4.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	// bestLines := make([]string, 500)
	var bestLines []string
    for scanner.Scan() {
		input := scanner.Text()
		
		bestLines = append(bestLines, createTableOfAllRepeatingKeyXORsWithInput(input))
    } 

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	fmt.Println(findMostEnglishString(bestLines))
}
func createTableOfAllRepeatingKeyXORsWithInput(input string) (bestString string){
	
	texts := make([]string, 255)
	decodedString := make([]string, 255)
	// fmt.Println("balalla")
	for letter := 0; letter < 255; letter++ {

		for i := 0; i < len(input)/len(hex.EncodeToString([]byte(string(letter)))); i++ {
			// fmt.Printf("\nletter: %v, i:%v\n", letter, i)
			texts[letter] += hex.EncodeToString([]byte(string(letter)))

		}
		//fmt.Println(len([]byte(texts[letter])))
		// fmt.Printf("contents of byte: %v\n",[]byte(texts[letter]))
		// fmt.Printf("contents of hex(byte):%s\n",texts[letter])
		// fmt.Println("")
		
		msg, err := challenge2.Fixed_XOR_on_hex_strings(input, texts[letter])
		if err != nil {
			fmt.Println(input, texts[letter])
			fmt.Println("PROBLEMS with XOR " + err.Error())
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

	return findMostEnglishString(decodedString)
}

func findMostEnglishString(decodedString []string) (bestString string) {
	var maxCount int
	for _ , element := range decodedString {
		currentCount := englishCount(element)
		if currentCount > maxCount {
			bestString = element
			maxCount = currentCount
		}
		// fmt.Println(currentCount)
	}
	return bestString

}