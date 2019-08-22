package main

import (
	"errors"
	// "../challenge2"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "strings"
	"math/bits"
)

func main() {
	fmt.Println(CumputeDistance("this is a test", "wokka wokka!!!"))

	file, err := os.Open("6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytesFromFile, err := ioutil.ReadAll(file)
	decodedBytesFromFile, err := base64.StdEncoding.DecodeString(string(bytesFromFile))

	fmt.Println(getNormalizedDistanceofKeysize(decodedBytesFromFile, 2))
}

func getNormalizedDistanceofKeysize(decodedBytesFromFile []byte, keysize int) float64 {

	slices := make([][]uint8, 4)       // initialize a slice of dy slices
	for i:=0;i<4;i++ {
		slices[i] = decodedBytesFromFile[i*keysize : keysize*(i+1)]
	}

	distanceSum := sumOfDistances(slices, 0)
	return float64(distanceSum) / float64(keysize) / 12
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
