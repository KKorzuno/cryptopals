package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"cryptopals/supportfunctions"
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
		inputBytes, err2 := supportfunctions.HexStringToBytes(input)
		if err2 != nil{
			log.Fatal(err2)
		}
		bestLines = append(bestLines, createTableOfAllRepeatingKeyXORsWithInput(string(inputBytes)))
    } 

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	fmt.Println(findMostEnglishString(bestLines))
}
func createTableOfAllRepeatingKeyXORsWithInput(input string) (bestString string){
	
	texts := make([]string, 128)
	decodedString := make([]string, 128)
	// fmt.Println("balalla")
	for letter := 0; letter < 128; letter++ {

		for i := 0; i < len(input)/len(string(letter)); i++ {
			// fmt.Printf("\nletter: %v, i:%v\n", letter, i)
			texts[letter] += string(letter)

		}
		//fmt.Println(len([]byte(texts[letter])))
		// fmt.Printf("contents of byte: %v\n",[]byte(texts[letter]))
		// fmt.Printf("contents of hex(byte):%s\n",texts[letter])
		// fmt.Println("")
		
		msg := supportfunctions.XOROnBytes([]byte(input), []byte(texts[letter]))

		decodedString[letter] = string(msg)
		//fmt.Println(decodedString[letter])
	}

	return findMostEnglishString(decodedString)
}

func findMostEnglishString(decodedString []string) (bestString string) {
	var maxCount int
	for _ , element := range decodedString {
		currentCount := supportfunctions.EnglishCount(element)
		if currentCount > maxCount {
			bestString = element
			maxCount = currentCount
		}
		// fmt.Println(currentCount)
	}
	return bestString

}