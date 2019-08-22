package main

import (
	"errors"
	"../challenge2"
	"../challenge5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"math/bits"
)

func main() {
	// fmt.Println(CumputeDistance("this is a test", "wokka wokka!!!"))

	file, err := os.Open("6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytesFromFile, err := ioutil.ReadAll(file)
	decodedBytesFromFile, err := base64.StdEncoding.DecodeString(string(bytesFromFile))

	distances := make([]float64, 40-2)
	for i:=3; i<=40;i++ {
		distances[i-3] = getNormalizedDistanceofKeysize(decodedBytesFromFile, i, 5)
	}
	
	//somehow guess the keysize
	keysize := 29

	
	// fmt.Println(len(decodedBytesFromFile))
	// for i,e := range distances {
	// 	fmt.Println( e, i+3)
	// }

	// for i:= 0; i< keysize: i++{

	// }
		// allTheLetters := make([]int, keysize)
		var allTheLetters string
	for i:=0;i<keysize;i++ {
		column := extractColumn(transposeBytes(decodedBytesFromFile,keysize),i)
		letterInt, _ := getMostProbableKeyLetter(hex.EncodeToString(column))
		allTheLetters = allTheLetters + string(byte(letterInt))
	}

	fmt.Println(allTheLetters)


	result, _ := challenge5.ApplyRepeatingKeyXOR(string(decodedBytesFromFile),allTheLetters)

	bytes, _ := hex.DecodeString(result)
	fmt.Println(string(bytes))

}

func getNormalizedDistanceofKeysize(decodedBytesFromFile []byte, keysize int, nSlices int) float64 {

	slices := make([][]uint8, nSlices)       // initialize a slice of dy slices
	for i:=0;i<nSlices;i++ {
		slices[i] = decodedBytesFromFile[i*keysize : keysize*(i+1)]
	}


	distance := sumOfDistances(slices,0)

	return float64(distance) / float64(keysize) / 12
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

	for i := 0; i < len(decoded_hex_string); i++ {
		result += bits.OnesCount(uint(decoded_hex_string[i]) ^ uint(decoded_hex_string2[i]))
	}
	return
}



func sumOfDistances(slice [][]byte, currentSum int) (sum int) {
	if len(slice) < 2 {return currentSum}
	sliceTail := slice[1:len(slice)]

		for _,e := range sliceTail {
			currentSum += ComputeHammingDistance(hex.EncodeToString(slice[0]), hex.EncodeToString(e))
		}
	
	return sumOfDistances(sliceTail, currentSum)
}

func sumOfDistances2(slices [][]byte) (distance int){
	for i,slice1 := range slices {
		for j,slice2 := range slices {
			if (i>=j) {continue}
			distance += ComputeHammingDistance(hex.EncodeToString(slice1), hex.EncodeToString(slice2))
		}
	}
	return distance
}

func transposeBytes(decodedBytesFromFile []byte, keysize int)( [][]byte ) {
	totalBytes := len(decodedBytesFromFile)
	transposedBytes := make([][]byte, totalBytes/keysize+1)
	
	for i:=0; i<totalBytes/keysize +1; i++ {
		transposedBytes[i] = decodedBytesFromFile[keysize * i:keysize * (i+1)]
	}
	return transposedBytes
}

func extractColumn(slice2d [][]byte, columnIndex int) (column []byte) {
    column = make([]byte, 0)
    for _, row := range slice2d {
		if(len(row) <= columnIndex) {
			break
		}
		column = append(column, row[columnIndex])
    }
    return
}

func getMostProbableKeyLetter(input string) (int,  string){
	
	repeatedLetters := make([]string, 255)
	decodedStrings := make([]string, 255)

	for letter := 0; letter < 255; letter++ {
		for i := 0; i < len(input)/len(hex.EncodeToString([]byte(string(letter)))); i++ {
			repeatedLetters[letter] += hex.EncodeToString([]byte(string(letter)))
		}
		// fmt.Println("HERE")
		// fmt.Println(len(repeatedLetters[letter]))
		// fmt.Println(len(input)) 
		// fmt.Println("HERE")

		msg, err := challenge2.Fixed_XOR_on_hex_strings(input, repeatedLetters[letter])
		if err != nil {
			// fmt.Println(input, repeatedLetters[letter])
			fmt.Println("PROBLEMS with XOR " + err.Error())
			return 0,""
		}
		str2, err2 := hex.DecodeString(msg)
		if err2 != nil {
			fmt.Println("PROBLEMS with decoding the outcome hex")
			// return 0,"" 
		}
		decodedStrings[letter] = string(str2)
	}

	return findMostEnglishString(decodedStrings)
}

func findMostEnglishString(decodedString []string) (letter int, bestString string) {
	var maxCount int
	for i , element := range decodedString {
		currentCount := englishCount(element)
		if currentCount > maxCount {
			letter = i
			bestString = element
			maxCount = currentCount
		}
	}
	return
}

func englishCount(input string) (count int) {
	input = strings.ToLower(input)
	for _ , element := range "etaoin shrdlu" {
		count += strings.Count(input, string(element))
	}
	return
}
