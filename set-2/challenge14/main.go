package main

import (
	//"cryptopals/set-2/challenge10"
	// "encoding/base64"
	"fmt"
	//"cryptopals/set-2/challenge11"
	//"cryptopals/set-1/challenge8"
	//"crypto/rand"
	//"math/big"
	"cryptopals/supportfunctions"

)



func main() {
	secretMessage := "NOBODY EXPECTS THE SPANISH INQUISITION!!!!"
	unknownKey := "YELLOW SUBMARINE"

	sillyOracle:= supportfunctions.NewRandomPrefixAttackerSecretECBOracle(string(secretMessage),unknownKey)

	keysize, _, _ := supportfunctions.DiscoverBlockSize(sillyOracle)
	fmt.Println(keysize)
	//Initialization for the first iteration, generating a keysize long list of A's
	// decryptedPaddedByte := supportfunctions.GetOneByteShortLetterList(keysize + 1)
	// var exists bool

	// encryptedBytes := sillyOracle.Encrypt(unknownString)

	// finalString := ""
	// for j := 0; j < len(encryptedBytes)/keysize; j++ {
	// 	for i := 0; i < keysize; i++ {
	// 		// fmt.Println(decryptedPaddedByte, len(decryptedPaddedByte))
	// 		// fmt.Println

	// 		encryptedBytes := encryptionOracle([]byte(getOneByteShortLetterList(keysize-i)), unknownString, unknownKey)
	// 		keysizeOfBytes := encryptedBytes[keysize*j:keysize*(j+1)]

	// 		cyphers := supportfunctions.getCypherMap(sillyOracle, unknownKey, keysize, decryptedPaddedByte[1:])

	// 		decryptedPaddedByte, exists = cyphers[string(keysizeOfBytes)]
	// 		if exists == false {
	// 			fmt.Println("FUCCK")
	// 			// fmt.Println(cyphers)
	// 		}
	// 	}
	// 	// fmt.Println(decryptedPaddedByte)
	// 	finalString = finalString + decryptedPaddedByte
	// }
	// fmt.Println(finalString)
}