package supportfunctions

import (
	"encoding/hex"
	"encoding/base64"
	"log"
	"strings"
	"math/bits"
	"fmt"
)

//XOROnBytes abc
func XOROnBytes(by1 []byte , by2 []byte ) []byte {

	if len(by1) != len(by2) {
		log.Fatal ("[]bytes to XOR not of equal lengths!")		
	}
	output:=make([]byte,len(by1))
	for i:=0; i<len(by1); i++ {
		output[i] = by1[i]^by2[i]
	}
	return output
}


//ComputeHammingDistance abc
func ComputeHammingDistance(by1 []byte, by2 []byte) (result int) {

	for index, bits1 := range by1 {
		bits2 := by2[index]
		result += bits.OnesCount(uint(bits1) ^ uint(bits2))
	}
	return result
}


//HexStringToBase64 abc
func HexStringToBase64(msg string) (output string, err error) {
	decodedHexString, err := hex.DecodeString(msg)
	output = base64.StdEncoding.EncodeToString(decodedHexString)
	//fmt.Println(output)
	return
}

//HexStringToBytes abc
func HexStringToBytes(msg string) (output []byte, err error) {
	output, err = hex.DecodeString(msg)
	return
}


//GetCypherMap abc
func GetCypherMap(sillyOracle Oracle, keysize int, oneByteShortByteList []byte) (map[string]string) {
	cyphers := make(map[string]string, 128)
	var temporaryString string
	for i := 0; i < 128; i++ {
		val := append(oneByteShortByteList, []byte(string(i))...)
		temp := sillyOracle.Encrypt([]byte(val))
		tempIn2D:= slicePaddedBytesInto2D(temp, keysize)
		key:=tempIn2D[0]
		temporaryString = string(key)
		// if(i==82) {temp82 = string(key)}
		// if(i==83) {temp83 = string(key)}
		cyphers[temporaryString] = string(val)
		// if(i>83) {
		// 	fmt.Println("KEY: ",temporaryString, "VAL: ",cyphers[temporaryString], i)
		// 	fmt.Println("KEY: ",temp82, "VAL: ",cyphers[temp82])
		// 	fmt.Println("KEY: ",temp83, "VAL: ",cyphers[temp83])
		//}

	}
	// fmt.Println("KEY: ",temporaryString, "VAL: ",cyphers[temporaryString])
	// fmt.Println("KEY: ",temporaryString, "VAL: ",cyphers[temporaryString])
	// fmt.Println("KEY: ",temp82, "VAL: ",cyphers[temp82])
	// fmt.Println("KEY: ",temp83, "VAL: ",cyphers[temp83])
	return cyphers
}

//GetByteListFromClonedString abc
func GetByteListFromClonedString(size int, letter string) []byte {
	var byteBunchOfBytes []byte
	for i := 0; i < size; i++ {
		byteBunchOfBytes = append(byteBunchOfBytes, []byte(letter)...)	
	}
	//fmt.Println("LEN OF PREPARED BYTE LIST:", len(byteBunchOfAs))
	return byteBunchOfBytes
}


//AddPadding Abc
func AddPadding(bytes []byte, keysize int) []byte{
	if len(bytes)%keysize == 0 {return bytes}
	//fmt.Println("doing appending")
	nPads := keysize - len(bytes)%keysize
	for i := 0; i<nPads; i++ {
		bytes = append(bytes ,[]byte("\x04")...)
	}
	//fmt.Println("len of byte inside padding function after padding: ", len(bytes))
	return bytes

}

//RemovePaddingAndFlatten2DArray Flattens and removes padding from a 2D array that it is provided
func RemovePaddingAndFlatten2DArray(bytesToFlatten [][]byte) (flatBytes []byte) {
	paddingSize:= strings.Count(string(bytesToFlatten[len(bytesToFlatten)-1]),"\x04") 
	for i,e := range bytesToFlatten {
		if (i == len(bytesToFlatten)-1){
			flatBytes = append(flatBytes, e[0:len(bytesToFlatten[i])-paddingSize]...)
			break
		}
		flatBytes = append(flatBytes, e...)

	}

	return
}
//GetNormalizedDistanceofKeysize DO NOT USE, OUTDATED
// func GetNormalizedDistanceofKeysize(decodedBytesFromFile []byte, keysize int, nSlices int) float64 {

// 	slices := make([][]uint8, nSlices)       // initialize a slice of dy slices
// 	for i:=0;i<nSlices;i++ {
// 		slices[i] = decodedBytesFromFile[i*keysize : keysize*(i+1)]
// 	}


// 	distance := sumOfDistances(slices,16)

// 	return float64(distance) / float64(keysize) / 12
// }

//GetMinDistanceInKeysizeMultiComparison Function to find the minimal Hamming distance for multiple ciphered blocks and if the distance is equal to 0,
//the position of the first doubled element
func GetMinDistanceInKeysizeMultiComparison(decodedBytesFromFile []byte, keysize int, nSlices int) (minDistanceFound int, positionofDoubleByte int) {

	slices := make([][]uint8, nSlices)       // initialize a slice of dy slices
	for i:=0;i<nSlices;i++ {
		slices[i] = decodedBytesFromFile[i*keysize : keysize*(i+1)]
	}
	//fmt.Println(slices)
	outsideValueForDoubleBytesPosition:=0
	return MinOfDistances(slices,1000, 0, &outsideValueForDoubleBytesPosition), outsideValueForDoubleBytesPosition
	
}

func sumOfDistances(slice [][]byte, currentSum int) (sum int) {
	if len(slice) < 2 {return currentSum}
	sliceTail := slice[1:len(slice)]

		for _,e := range sliceTail {
			currentSum += ComputeHammingDistance(slice[0], e)
		}
	
	return sumOfDistances(sliceTail, currentSum)
}

//MinOfDistances Recursive function to go through a 2D slice of bytes to calculate a minimal distance and if 
func MinOfDistances(slice [][]byte, currentMin int, tempSliceNumber int, outsideValueForDoubleBytesPosition *int) (sum int) {
	if len(slice) < 2 {return currentMin}
	sliceTail := slice[1:len(slice)]
		for _,e := range sliceTail {
			fmt.Printf("Current minimum : %v\n",currentMin)
			temp := ComputeHammingDistance(slice[0], e)
			if currentMin > temp {
				currentMin = temp
				fmt.Printf("CURRENTMIN > TEMP, Current minimum : %v and current temp: %v\n",currentMin,temp)
				if temp==0 {
					*outsideValueForDoubleBytesPosition = tempSliceNumber
					fmt.Println("IM IN DOUBLE BYTE ASSIGNMENT")
				}
			} 
		}
		tempSliceNumber++
	return MinOfDistances(sliceTail, currentMin, tempSliceNumber, outsideValueForDoubleBytesPosition)
}

//CheckIfMinDistanceIsEqual3Times UNTESTED, DO NOT USE
func CheckIfMinDistanceIsEqual3Times(decodedBytesFromFile []byte, keysize int) bool {

	slices := make([][]uint8, len(decodedBytesFromFile)/keysize)       
	for i:=0;i<len(decodedBytesFromFile)/keysize;i++ {
		slices[i] = decodedBytesFromFile[i*keysize : keysize*(i+1)]
	}

	return FindThreeEqualDistances(slices)
}
//FindThreeEqualDistances UNTESTED, DO NOT USE
func FindThreeEqualDistances(slice [][]byte) (ThreeEqual bool) {
	if len(slice) < 3 {return false}
	sliceTail := slice[1:len(slice)]
	ThreeEqual = false
	FoundEqualDistanceOnce := false
		for _,e := range sliceTail {
			temp := ComputeHammingDistance(slice[0], e)
			if temp == 0 &&	FoundEqualDistanceOnce == true {
				 return true
			} else if temp == 0 {
				FoundEqualDistanceOnce = true
			}
		}
		return FindThreeEqualDistances(sliceTail)
}