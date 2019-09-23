package supportfunctions

import (
	"crypto/rand"
	"fmt"
	"math/big"
)


//GetRandomBytes abc
func GetRandomBytes(numberOfBytes int) (randomBytes []byte) {

	randomBytes = make([]byte, numberOfBytes)
	_, err := rand.Read(randomBytes)
	//fmt.Println(n, err, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	return randomBytes
}

//GetRandomInt abc
func GetRandomInt(maxValue int) (int, error) {
	randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(maxValue)))
	return int(randomInt.Int64()), err
}

//AppendRandomBytesInFrontAndBack abc
func AppendRandomBytesInFrontAndBack(input []byte)(output []byte){
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