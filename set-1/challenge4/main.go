package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"cryptopals/supportfunctions"
)



func main() {
	//TO DO: MODIFY YOUR PATH!!
	file, err := os.Open("/home/kkorzuno/go/src/cryptopals/set-1/challenge4/4.txt")
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
		TableOfRepeatingLettersXORedwithInput := supportfunctions.CreateTableOfAllRepeatingKeyXORsWithInput(string(inputBytes))
		_ , singleGoodString:= supportfunctions.FindMostEnglishString(TableOfRepeatingLettersXORedwithInput)
		bestLines = append(bestLines,singleGoodString)
    } 

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	_, bestString:= supportfunctions.FindMostEnglishString(bestLines)
	fmt.Println(bestString)

}


