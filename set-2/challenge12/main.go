package challenge12

import (
	"cryptopals/set-2/challenge10"
	"encoding/base64"
	"fmt"
	// "cryptopals/set-2/challenge12"
	"cryptopals/set-1/challenge8"
)

func main() {
	unknownStringbase64 := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
	aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
	dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
	YnkK`
	unknownString, _ := base64.StdEncoding.DecodeString(unknownStringbase64)
	unknownKey := "YELLOW SUBMARINE"

	keysize := DiscoverBlocksize(unknownString, unknownKey)

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
	// fmt.Println(string(append(input,unknownString...)))
	encryptedByte = challenge10.EncryptEBC(unknownKey, append(input, unknownString...))

	return
}

func DiscoverBlocksize(unknownString []byte, unknownKey string) (blockSize int) {
	listOfAs := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	for i := 1; i < 15; i++ {
		listOfAs = listOfAs + "A"
		minDistance := challenge8.GetMinDistanceInKeysizeMultiComparison(encryptionOracle([]byte(listOfAs), unknownString, unknownKey), i, 2)

		if minDistance == 0 {
			return i
		} else {
			listOfAs := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
			for i := 1; i < 15; i++ {
				listOfAs = listOfAs + "A"
				//minDistance := challenge8.GetMinDistanceInKeysizeMultiComparison(encryptionOracle([]byte(listOfAs), unknownString, unknownKey), i, 3)

			}
		}
	}
	return 0
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
