package challenge8

import (

	//"cryptopals/supportfunctions"

	//"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	//"strings"
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

	//keysize:=16

	scanner := bufio.NewScanner(file)
	//var distances []float64
	var iteration int
	for scanner.Scan() {
		//input := scanner.Text()
		//tempBytes, _ := base64.StdEncoding.DecodeString(input)
		//distances = append(distances, getNormalizedDistanceofKeysize(tempBytes, keysize, 5))
		//fmt.Println(GetMinDistanceInKeysizeMultiComparison(tempBytes, keysize, 15), iteration)
		iteration++
	}
}

