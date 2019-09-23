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
func GetRandomInt(maxValue int) int {
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(maxValue)))
	return int(randomInt.Int64())
}
