package supportfunctions

import (
	//"cryptopals/set-2/challenge10"
	//"errors"
	"crypto/cipher"
	"crypto/aes"
	"os"
	"fmt"
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

//PrefixInputPostfixCBCOracle aaa
type PrefixInputPostfixCBCOracle struct {
	stringToPrefix string
	stringToPostfix string
	keysize int
	mycipher cipher.Block
}

//RandomPrefixInputPostfixECBOOrCBCOracle abc
type RandomPrefixInputPostfixECBOOrCBCOracle struct {
	keysize int
	CBCOracle Oracle
	EBCOracle Oracle
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

//Decrypt abc
func (myOracle *PrefixInputPostfixECBOracle) Decrypt(bytesToDecrypt []byte) ([]byte) {

	
	bytesToDecryptIn2D := slicePaddedBytesInto2D(bytesToDecrypt,myOracle.keysize)
	DecryptedBytesIn2D := slicePaddedBytesInto2D(bytesToDecrypt,myOracle.keysize)
	for i := 0; i < len(bytesToDecryptIn2D); i++ {
		myOracle.mycipher.Encrypt(DecryptedBytesIn2D[i], bytesToDecryptIn2D[i])
	}

	return RemovePaddingAndFlatten2DArray(DecryptedBytesIn2D)
}


//NewPrefixInputPostfixCBCOracle abc
func NewPrefixInputPostfixCBCOracle(stringToPrefix string, stringToPostfix string, secretKey string) Oracle {
	aesCipher, err :=  aes.NewCipher([]byte(secretKey))
	if err != nil {
		fmt.Println("Wrong input to cipher creation")
		os.Exit(1)
	}
	return &PrefixInputPostfixECBOracle{stringToPrefix, stringToPostfix, len(secretKey), aesCipher}
}

//PrepareInputString abc
func (myOracle *PrefixInputPostfixCBCOracle) PrepareInputString(input string)(output string) {

	output = myOracle.stringToPrefix + input + myOracle.stringToPostfix 
	//fmt.Println("PREPARE INPUT STRING:", output)
	output = string(AddPadding([]byte(output), myOracle.keysize))
	//fmt.Println("Prepared Input String for Oracle",(output))

	return
}

//Encrypt abc
func (myOracle *PrefixInputPostfixCBCOracle) Encrypt(unpaddedBytesToEncrypt []byte, iv []byte) ([]byte) {
	
	

	paddedStringToEncrypt := myOracle.PrepareInputString(string(unpaddedBytesToEncrypt))
	supportVector := []byte (iv)
	bytesToEncryptIn2D := make([][]byte, len(paddedStringToEncrypt)/myOracle.keysize)
	encryptedBytesIn2D := make([][]byte, len(paddedStringToEncrypt)/myOracle.keysize)


	for i := 0; i < len(bytesToEncryptIn2D); i++ {
		bytesToEncryptIn2D[i] = []byte(paddedStringToEncrypt[myOracle.keysize*i : myOracle.keysize*(i+1)])
		bytesToEncryptIn2D[i] = XOROnBytes(bytesToEncryptIn2D[i], supportVector)
		encryptedBytesIn2D[i] = make([]byte, myOracle.keysize)
		// fmt.Println(encryptedBytesIn2D[i], bytesToEncryptIn2D[i])
		//mycipher.Encrypt(encryptedBytesIn2D[i], bytesToEncryptIn2D[i])
		supportVector=encryptedBytesIn2D[i]
	}

	var encryptedBytesFlat []byte
	for _,e := range encryptedBytesIn2D {
		encryptedBytesFlat = append(encryptedBytesFlat, e...)
	}

	return encryptedBytesFlat
	
	
	
	
	
	
	
	// //fmt.Println(unpaddedBytesToEncrypt)
	// paddedStringToEncrypt := myOracle.PrepareInputString(string(unpaddedBytesToEncrypt))
	// //fmt.Println("ENCRYPT: ", paddedStringToEncrypt, "BYTES: ", []byte(paddedStringToEncrypt))
	// bytesToEncryptIn2D := slicePaddedBytesInto2D([]byte(paddedStringToEncrypt),myOracle.keysize)
	// encryptedBytesIn2D := slicePaddedBytesInto2D([]byte(paddedStringToEncrypt),myOracle.keysize)
	// for i := 0; i < len(bytesToEncryptIn2D); i++ {
	// 	myOracle.mycipher.Encrypt(encryptedBytesIn2D[i], bytesToEncryptIn2D[i])
	// }
 	// var encryptedBytesFlat []byte
	// for _,e := range encryptedBytesIn2D {
	// 	encryptedBytesFlat = append(encryptedBytesFlat, e...)
	// }
	// return encryptedBytesFlat
}

//Decrypt abc
func (myOracle *PrefixInputPostfixCBCOracle) Decrypt(bytesToDecrypt []byte) ([]byte) {

	
	bytesToDecryptIn2D := slicePaddedBytesInto2D(bytesToDecrypt,myOracle.keysize)
	DecryptedBytesIn2D := slicePaddedBytesInto2D(bytesToDecrypt,myOracle.keysize)
	for i := 0; i < len(bytesToDecryptIn2D); i++ {
		myOracle.mycipher.Encrypt(DecryptedBytesIn2D[i], bytesToDecryptIn2D[i])
	}

	return RemovePaddingAndFlatten2DArray(DecryptedBytesIn2D)
}

func slicePaddedBytesInto2D (inputByte []byte, lenghtOfaRow int) (slicesIn2D [][]byte){
	slicesIn2D= make([][]byte, len(inputByte)/lenghtOfaRow)       
	for i:=0;i<len(inputByte)/lenghtOfaRow;i++ {
		slicesIn2D[i] = inputByte[i*lenghtOfaRow : lenghtOfaRow*(i+1)]
	}
	return
}





