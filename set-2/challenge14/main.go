package main

import (
	"cryptopals/set-2/challenge10"
	// "encoding/base64"
	"fmt"
	"cryptopals/set-2/challenge11"
	"cryptopals/set-1/challenge8"
	"crypto/rand"
	"math/big"
)

func main() {

	unknownString := []byte("Hello world, this beautiful world :) The sun is shining")
	unknownKey := "YELLOW SUBMARINE"

	keysize := discoverBlocksize(unknownString, unknownKey)

	//Initialization for the first iteration, generating a keysize long list of A's
	decryptedPaddedByte := getOneByteShortLetterList(keysize + 1)
	var exists bool

	encryptedBytes := encryptionOracle([]byte(""), unknownString, unknownKey)

	finalString := ""
	for j := 0; j < len(encryptedBytes)/keysize; j++ {
		for i := 0; i < keysize; i++ {
			// fmt.Println(decryptedPaddedByte, len(decryptedPaddedByte))
			// fmt.Println

			encryptedBytes := encryptionOracle([]byte(getOneByteShortLetterList(keysize-i)), unknownString, unknownKey)
			keysizeOfBytes := encryptedBytes[keysize*j:keysize*(j+1)]

			cyphers := getCypherMap(unknownKey, keysize, decryptedPaddedByte[1:])

			decryptedPaddedByte, exists = cyphers[string(keysizeOfBytes)]
			if exists == false {
				fmt.Println("FUCCK")
				// fmt.Println(cyphers)
			}
		}
		// fmt.Println(decryptedPaddedByte)
		finalString = finalString + decryptedPaddedByte
	}
	fmt.Println(finalString)
}

//AAAAAAAAAAAAAARBF
func encryptionOracle(input []byte, unknownString []byte, unknownKey string) (encryptedByte []byte) {
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(100))
	randomPrefix := challenge11.GetRandomBytes(int(randomInt.Int64()))
	encryptedByte = challenge10.EncryptEBC(unknownKey, append(randomPrefix, append(input, unknownString...)...))

	return
}




func getCypherMap(unknownKey string, keysize int, oneByteShortLetterList string) map[string]string {
	cyphers := make(map[string]string, 255)

	for i := 0; i < 255; i++ {
		val := oneByteShortLetterList + string(i)
		// key:= string(challenge10.EncryptEBC(unknownKey,[]byte(val)))
		key := string(encryptionOracle([]byte(""), []byte(val), unknownKey))

		cyphers[key] = val
	}

	return cyphers
}

func getOneByteShortLetterList(keysize int) string {
	bunchOfAs := ""
	for i := 0; i < keysize-1; i++ {
		bunchOfAs = bunchOfAs + "A"
	}
	return bunchOfAs
}
