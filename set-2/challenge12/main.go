package main

import (
	"fmt"
	"encoding/base64"
	"cryptopals/set-2/challenge10"
	// "cryptopals/set-2/challenge12"
	"cryptopals/set-1/challenge8"
	

)

func main () {
	unknownStringbase64 := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
	aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
	dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
	YnkK`
	unknownString, _ := base64.StdEncoding.DecodeString(unknownStringbase64)
	unknownKey := "YELLOW SUBMARINE"

// encryptionOracle(unknownString,[]byte("asdasd"),unknownKey)
	keysize := discoverBlocksize(unknownString,unknownKey)

	cyphers := getCypherMap(unknownKey, keysize)

	
}


func encryptionOracle(unknownString []byte, input []byte, unknownKey string) (encryptedByte []byte) {
		encryptedByte = challenge10.EncryptEBC(unknownKey,append(input,unknownString...))
	
	return 
}

func discoverBlocksize (unknownString []byte,unknownKey string) (blockSize int) {
		listOfAs := ""
	for i:=1;i<64;i++ {
		listOfAs = listOfAs + "AA"
		minDistance := challenge8.GetMinDistanceInKeysizeMultiComparison(encryptionOracle(unknownString,[]byte(listOfAs),unknownKey), i, 2)

		if(minDistance == 0) {
			return i
		}
	}
	return 0
}

func getCypherMap(unknownKey string, keysize int) map[string]string {

	bunchOfAs := getOneByteShortLetterList(keysize)
	cyphers := make(map[string]string,255)

	for i:=0;i<255;i++ {
		val:= bunchOfAs + string(i)
		key:= string(challenge10.EncryptEBC(unknownKey,[]byte(val)))
		cyphers[key] = val
		fmt.Println(key, val)
	}

	return cyphers
}

func getOneByteShortLetterList(keysize int) string{
	bunchOfAs := ""
	for i:=0;i<keysize-1;i++ {
		bunchOfAs = bunchOfAs + "A"
	}
	return bunchOfAs
}