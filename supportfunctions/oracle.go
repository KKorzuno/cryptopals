package supportfunctions


import (
	//"cryptopals/set-2/challenge10"
	//"errors"
	"crypto/cipher"
	"crypto/aes"
	"strings"
	"os"
	"fmt"
	"cryptopals/set-1/challenge8"
)

//Oracle interface type for all encryption functions
type Oracle interface {
	PrepareInputString(input string) (output string)
	Encrypt([]byte) ([]byte)
	Decrypt([]byte) ([]byte)
}
//PrefixInputPostfixECBOracle aaa
type PrefixInputPostfixECBOracle struct {
	stringToPrefix string
	stringToPostfix string
	keysize int
	mycipher cipher.Block
}

//NewPrefixInputPostfixECBOracle abc
func NewPrefixInputPostfixECBOracle(stringToPrefix string, stringToPostfix string, secretKey string) Oracle {
	aesCipher, err :=  aes.NewCipher([]byte(secretKey))
	if err != nil {
		fmt.Println("Wrong input to cipher creation")
		os.Exit(1)
	}
	return &PrefixInputPostfixECBOracle{stringToPrefix, stringToPostfix, len(secretKey), aesCipher}
}

//PrepareInputString abc
func (myOracle *PrefixInputPostfixECBOracle) PrepareInputString(input string)(output string) {

	output = myOracle.stringToPrefix + input + myOracle.stringToPostfix 
	//fmt.Println("PREPARE INPUT STRING:", output)
	output = string(AddPadding([]byte(output), myOracle.keysize))
	//fmt.Println("Prepared Input String for Oracle",(output))

	return
}

//Encrypt abc
func (myOracle *PrefixInputPostfixECBOracle) Encrypt(unpaddedBytesToEncrypt []byte) ([]byte) {
	//fmt.Println(unpaddedBytesToEncrypt)
	paddedStringToEncrypt := myOracle.PrepareInputString(string(unpaddedBytesToEncrypt))
	//fmt.Println("ENCRYPT: ", paddedStringToEncrypt, "BYTES: ", []byte(paddedStringToEncrypt))
	bytesToEncryptIn2D := slicePaddedBytesInto2D([]byte(paddedStringToEncrypt),myOracle.keysize)
	encryptedBytesIn2D := slicePaddedBytesInto2D([]byte(paddedStringToEncrypt),myOracle.keysize)
	for i := 0; i < len(bytesToEncryptIn2D); i++ {
		myOracle.mycipher.Encrypt(encryptedBytesIn2D[i], bytesToEncryptIn2D[i])
	}
 	var encryptedBytesFlat []byte
	for _,e := range encryptedBytesIn2D {
		encryptedBytesFlat = append(encryptedBytesFlat, e...)
	}
	return encryptedBytesFlat
}

func slicePaddedBytesInto2D (inputByte []byte, lenghtOfaRow int) (slicesIn2D [][]byte){
	slicesIn2D= make([][]byte, len(inputByte)/lenghtOfaRow)       
	for i:=0;i<len(inputByte)/lenghtOfaRow;i++ {
		slicesIn2D[i] = inputByte[i*lenghtOfaRow : lenghtOfaRow*(i+1)]
	}
	return
}



// func (myOracle *PrefixInputPostfixECBOracle) Encrypt() ([]byte, error) {
	
// 	if myOracle.stringToEncrypt == "" {
// 		return []byte(""), errors.New("Wrong Input") 
// 	}
// 	return challenge10.EncryptEBC(myOracle.secretKey, []byte(myOracle.stringToEncrypt)), nil
// }

//Decrypt abc
func (myOracle *PrefixInputPostfixECBOracle) Decrypt(bytesToDecrypt []byte) ([]byte) {

	
	bytesToDecryptIn2D := slicePaddedBytesInto2D(bytesToDecrypt,myOracle.keysize)
	DecryptedBytesIn2D := slicePaddedBytesInto2D(bytesToDecrypt,myOracle.keysize)
	for i := 0; i < len(bytesToDecryptIn2D); i++ {
		myOracle.mycipher.Encrypt(DecryptedBytesIn2D[i], bytesToDecryptIn2D[i])
	}

	return RemovePaddingAndFlatten2DArray(DecryptedBytesIn2D)
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

//DiscoverBlockSize abc
func DiscoverBlockSize(sillyOracle Oracle) (estimatedKeyLength int, addedPseudoPadding int, doubleBytePosition int) {
	listOfAs := strings.Repeat("A", 16*2)
	estimatedKeyLength = 16
	for addedPseudoPadding = 0; addedPseudoPadding < 16; addedPseudoPadding++ {
		
		encryptedProfileWithAs := sillyOracle.Encrypt([]byte(listOfAs))
		minDistance := -1
		minDistance, doubleBytePosition = challenge8.GetMinDistanceInKeysizeMultiComparison(encryptedProfileWithAs, estimatedKeyLength, len(encryptedProfileWithAs)/estimatedKeyLength )

		if minDistance == 0 {
			estimatedKeyLength = 16
			return
			
		}
		listOfAs = listOfAs + "A"
	}

	//UNTESTED CODE FOR BOTH 24 and 32 CASES
	estimatedKeyLength = 24
	fmt.Println("AFTER 16 BYTE CHECK")
	listOfAs = strings.Repeat("A", 24*3)
	for addedPseudoPadding = 0; addedPseudoPadding < 24; addedPseudoPadding++ {
		listOfAs = listOfAs + "A"
		encryptedProfileWithAs := sillyOracle.Encrypt([]byte(listOfAs))
		if challenge8.CheckIfMinDistanceIsEqual3Times(encryptedProfileWithAs, estimatedKeyLength) {
			estimatedKeyLength = 24
			return
		}
	}
	estimatedKeyLength = 32
	fmt.Println("AFTER 24 BYTE CHECK")
	listOfAs = strings.Repeat("A", 32*3)
	for addedPseudoPadding = 0; addedPseudoPadding < 32; addedPseudoPadding++ {
		listOfAs = listOfAs + "A"
		encryptedProfileWithAs := sillyOracle.Encrypt([]byte(listOfAs))
		if challenge8.CheckIfMinDistanceIsEqual3Times(encryptedProfileWithAs, estimatedKeyLength) {
			return 
		}
	}
	fmt.Println("AFTER 32 BYTE CHECK")
	return -1, 0, -1
}

//GetCypherMap abc
func GetCypherMap(sillyOracle Oracle, keysize int, oneByteShortByteList []byte) (map[string]string) {
	cyphers := make(map[string]string, 128)
	var temporaryString string
	//var temp82,temp83 string
	for i := 0; i < 128; i++ {
		//fmt.Println("oneByteShortByteList", string(oneByteShortByteList))
		val := append(oneByteShortByteList, []byte(string(i))...)
		//fmt.Println("LENGHT OF VAL AFTER append: ", len(val))
		//fmt.Println([]byte(val))
		//key := string(sillyOracle.Encrypt([]byte(val)))
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

//GetOneByteShortByteList abc
func GetOneByteShortByteList(size int) []byte {
	var byteBunchOfAs []byte
	// -2 due to size given at 16, m
	for i := 0; i < size-1; i++ {
		byteBunchOfAs = append(byteBunchOfAs, []byte("A")...)	
	}
	//fmt.Println("LEN OF PREPARED BYTE LIST:", len(byteBunchOfAs))
	return byteBunchOfAs
}
//PaddingOracleAttack abc
func PaddingOracleAttack (attackedOracle Oracle)(textBehindInput string){
	
	keysize, _, _ := DiscoverBlockSize(attackedOracle)
	fmt.Println(keysize)
	//Initialization for the first iteration, generating a keysize long list of A's to measure total number of bytes
	decryptedLetterPlusPadding := GetOneByteShortByteList(keysize + 1)
	//Passing empty just to see how long the output is, to calculate the number of iterations to scan through all
	encryptedBytes := attackedOracle.Encrypt([]byte(""))
	
	fmt.Println("MAIN LOOP STARTS")
	
	finalString := ""
	for j := 0; j < len(encryptedBytes)/keysize; j++ {
	
		for i := 0; i < keysize; i++ {
			
			attackerString := GetOneByteShortByteList(keysize - i)
			//fmt.Println("attackerString", attackerString)
			encryptedBytes := attackedOracle.Encrypt([]byte(attackerString))
			keysizeOfBytes := encryptedBytes[keysize*j:keysize*(j+1)]
			//fmt.Println("decryptedLetterPlusPadding",decryptedLetterPlusPadding, string(decryptedLetterPlusPadding))
			//fmt.Println("ENCRYPTED KEYSIZE of BYTES: ", keysizeOfBytes, "j:", j, "i:", i)
			//fmt.Println("JESTEM")
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