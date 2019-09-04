package challenge8

import (

	"cryptopals/set-1/challenge2"

	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"math/bits"
	"bufio"
)

func main() {
	// fmt.Println(CumputeDistance("this is a test", "wokka wokka!!!"))

	file, err := os.Open("8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	keysize:=16

	scanner := bufio.NewScanner(file)
	//var distances []float64
	var iteration int
	for scanner.Scan() {
		input := scanner.Text()
		tempBytes, _ := base64.StdEncoding.DecodeString(input)
		//distances = append(distances, getNormalizedDistanceofKeysize(tempBytes, keysize, 5))
		fmt.Println(GetMinDistanceInKeysizeMultiComparison(tempBytes, keysize, 15), float64(iteration))
		iteration++
	}
}

func GetNormalizedDistanceofKeysize(decodedBytesFromFile []byte, keysize int, nSlices int) float64 {

	slices := make([][]uint8, nSlices)       // initialize a slice of dy slices
	for i:=0;i<nSlices;i++ {
		slices[i] = decodedBytesFromFile[i*keysize : keysize*(i+1)]
	}


	distance := sumOfDistances(slices,16)

	return float64(distance) / float64(keysize) / 12
}

func GetMinDistanceInKeysizeMultiComparison(decodedBytesFromFile []byte, keysize int, nSlices int) float64 {

	slices := make([][]uint8, nSlices)       // initialize a slice of dy slices
	for i:=0;i<nSlices;i++ {
		slices[i] = decodedBytesFromFile[i*keysize : keysize*(i+1)]
	}

	return float64(MinOfDistances(slices,1000))
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


func sumOfDistances(slice [][]byte, currentSum int) (sum int) {
	if len(slice) < 2 {return currentSum}
	sliceTail := slice[1:len(slice)]

		for _,e := range sliceTail {
			currentSum += ComputeHammingDistance(hex.EncodeToString(slice[0]), hex.EncodeToString(e))
		}
	
	return sumOfDistances(sliceTail, currentSum)
}

func MinOfDistances(slice [][]byte, currentMin int) (sum int) {
	if len(slice) < 2 {return currentMin}
	sliceTail := slice[1:len(slice)]

		for _,e := range sliceTail {
			temp := ComputeHammingDistance(hex.EncodeToString(slice[0]), hex.EncodeToString(e))
			if currentMin > temp {
				currentMin = temp
			} 
		}
	
	return MinOfDistances(sliceTail, currentMin)
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