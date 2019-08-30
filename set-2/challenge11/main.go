package main

import (
	//"crypto/aes"
	// "encoding/base64"
	"fmt"
	"log"
	"os"
	// "encoding/hex"
	// "encoding/base64"
	"io/ioutil"
	// "cryptopals/set-1/challenge2"
	"cryptopals/set-2/challenge10"
	"cryptopals/set-1/challenge8"
	"crypto/rand"
	"math/big"
)



func main(){
	
	file, err := os.Open("song_short.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// bytesFromFile, err := ioutil.ReadAll(file)
	input, err := ioutil.ReadAll(file)

	fmt.Println(challenge8.GetNormalizedDistanceofKeysize(encryptionOracle(input),16,5))

	// fmt.Println(encryptionOracle(input))
	// fmt.Println("***********************")

	

}

func GetRandomBytes(numberOfBytes int)(randomBytes []byte){

	randomBytes = make([]byte, numberOfBytes)
	_, err := rand.Read(randomBytes)
	//fmt.Println(n, err, b)
	if err != nil{
		fmt.Println(err.Error())
	}
	return randomBytes
}

func encryptionOracle(input []byte) (encryptedByte []byte){
	flagInt, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(input)
	//fmt.Println("_________________")
	input = appendRandomBytesInFrontAndBack(input)
	//fmt.Println(input)

	if int(flagInt.Int64()) == 0 {
		fmt.Println("encrypting with CBC")
		encryptedByte = challenge10.EncryptCBC("abcdefghijklmnop",string(GetRandomBytes(16)),input)
	} else {
		fmt.Println("encrypting with EBC")
		encryptedByte = challenge10.EncryptEBC(string(GetRandomBytes(16)),input)
	}
	
	return 
}

func appendRandomBytesInFrontAndBack(input []byte)(output []byte){
	zeroToFive, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		fmt.Println(err.Error())
	}
	randomNumberFiveToTen := 5 + int(zeroToFive.Int64())
	tempBytes := GetRandomBytes(randomNumberFiveToTen)
	//fmt.Printf("\ntempBytes: %v\n" ,tempBytes)	
	output = append(tempBytes,input...)
	output = append(output, tempBytes...)

	return
}