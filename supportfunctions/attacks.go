package supportfunctions

import (
	"fmt"
	"strings"

)

//PaddingOracleAttack abc
func PaddingOracleAttack (attackedOracle Oracle)(string){
	
	keysize, _, _ := DiscoverBlockSize(attackedOracle)
	fmt.Println(keysize)
	//Initialization for the first iteration, generating a keysize long list of A's to measure total number of bytes
	decryptedLetterPlusPadding := GetByteListFromClonedString(keysize, "A")
	//Passing empty just to see how long the output is, to calculate the number of iterations to scan through all
	encryptedBytes := attackedOracle.Encrypt([]byte(""))
	
	fmt.Println("MAIN LOOP STARTS")
	
	finalString := ""
	for j := 0; j < len(encryptedBytes)/keysize; j++ {
	
		for i := 0; i < keysize; i++ {
			
			attackerString := GetByteListFromClonedString(keysize -1 - i,"A")
			//fmt.Println("attackerString", attackerString)
			encryptedBytes := attackedOracle.Encrypt([]byte(attackerString))
			keysizeOfBytes := encryptedBytes[keysize*j:keysize*(j+1)]
			//fmt.Println("decryptedLetterPlusPadding",decryptedLetterPlusPadding, string(decryptedLetterPlusPadding))
			//fmt.Println("ENCRYPTED KEYSIZE of BYTES: ", keysizeOfBytes, "j:", j, "i:", i)
			cyphers := GetCypherMap(attackedOracle, keysize, decryptedLetterPlusPadding[1:])
			// for x,s := range cyphers{
			// 	fmt.Println("ELEMENT CYPHERS:", x, s)
			// }
			//fmt.Println("MAP SIZE", len(cyphers))
			//fmt.Println(cyphers)
			//decryptedLetterPlusPadding, exists = cyphers[string(keysizeOfBytes)]
			tempString, exists := cyphers[string(keysizeOfBytes)]
			if exists == false {
				fmt.Println("DIDNT FIND ELEMENT IN THE CYPHERMAP")
			}
			decryptedLetterPlusPadding = []byte(tempString)
			//fmt.Println("decryptedLetterPlusPadding",decryptedLetterPlusPadding, len(decryptedLetterPlusPadding))

			
		}
		//fmt.Println("DECRYPTED LETTER PLUS PADDING:" ,decryptedLetterPlusPadding)
		finalString = finalString + string(decryptedLetterPlusPadding)
		//fmt.Println(finalString)
	}
	return finalString
}

func getMostProbableKeyLetter(input string) (int,  string){
	
	repeatedLetters := make([]string, 255)
	decodedStrings := make([]string, 255)

	for letter := 0; letter < 255; letter++ {
		for i := 0; i < len(input)/len(string(letter)); i++ {
			repeatedLetters[letter] += string(letter)
		}
		// fmt.Println("HERE")
		// fmt.Println(len(repeatedLetters[letter]))
		// fmt.Println(len(input)) 
		// fmt.Println("HERE")

		msg := XOROnBytes([]byte(input), []byte(repeatedLetters[letter]))

		decodedStrings[letter] = string(msg)
	}

	return findMostEnglishString(decodedStrings)
}

func findMostEnglishString(decodedString []string) (letter int, bestString string) {
	var maxCount int
	for i , element := range decodedString {
		currentCount := EnglishCount(element)
		if currentCount > maxCount {
			letter = i
			bestString = element
			maxCount = currentCount
		}
	}
	return
}

//EnglishCount Counts all the instances of letters in "etaoin shrdlu"
func EnglishCount(input string) (count int) {
	input = strings.ToLower(input)
	for _ , element := range "etaoin shrdlu" {
		count += strings.Count(input, string(element))
	}
	return
}

//DiscoverBlockSize abc
func DiscoverBlockSize(sillyOracle Oracle) (estimatedKeyLength int, addedPseudoPadding int, doubleBytePosition int) {
	listOfAs := strings.Repeat("A", 16*2)
	estimatedKeyLength = 16
	for addedPseudoPadding = 0; addedPseudoPadding < 16; addedPseudoPadding++ {
		
		encryptedProfileWithAs := sillyOracle.Encrypt([]byte(listOfAs))
		minDistance := -1
		minDistance, doubleBytePosition = GetMinDistanceInKeysizeMultiComparison(encryptedProfileWithAs, estimatedKeyLength, len(encryptedProfileWithAs)/estimatedKeyLength )

		if minDistance == 0 {
			estimatedKeyLength = 16
			return
			
		}
		listOfAs = listOfAs + "A"
	}

	// //UNTESTED CODE FOR BOTH 24 and 32 CASES
	// estimatedKeyLength = 24
	// fmt.Println("AFTER 16 BYTE CHECK")
	// listOfAs = strings.Repeat("A", 24*3)
	// for addedPseudoPadding = 0; addedPseudoPadding < 24; addedPseudoPadding++ {
	// 	listOfAs = listOfAs + "A"
	// 	encryptedProfileWithAs := sillyOracle.Encrypt([]byte(listOfAs))
	// 	if challenge8.CheckIfMinDistanceIsEqual3Times(encryptedProfileWithAs, estimatedKeyLength) {
	// 		estimatedKeyLength = 24
	// 		return
	// 	}
	// }
	// estimatedKeyLength = 32
	// fmt.Println("AFTER 24 BYTE CHECK")
	// listOfAs = strings.Repeat("A", 32*3)
	// for addedPseudoPadding = 0; addedPseudoPadding < 32; addedPseudoPadding++ {
	// 	listOfAs = listOfAs + "A"
	// 	encryptedProfileWithAs := sillyOracle.Encrypt([]byte(listOfAs))
	// 	if challenge8.CheckIfMinDistanceIsEqual3Times(encryptedProfileWithAs, estimatedKeyLength) {
	// 		return 
	// 	}
	// }
	// fmt.Println("AFTER 32 BYTE CHECK")
	return -1, 0, -1
}
