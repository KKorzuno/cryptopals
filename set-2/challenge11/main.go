package main

import (

	"fmt"

	"cryptopals/set-2/challenge10"
	"cryptopals/supportfunctions"
	"crypto/rand"
	"math/big"
)



func main(){
	

	//fmt.Println(len(input))
	//cyphertext := encryptionOracle(input)
	//fmt.Println(len(cyphertext))
	if(0 == challenge8.GetMinDistanceInKeysizeMultiComparison(encryptionOracle([]byte("YELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINE")),16,5)){
		fmt.Println("Encrypted with ECB")
	} else {
	fmt.Println("Encrypted with CBC")
	}	
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
	//fmt.Println("len before appending with random bytes: ", len(input))
	input = appendRandomBytesInFrontAndBack(input)
	//fmt.Println(input)

	//fmt.Println("len after appending with random bytes: ", len(input))



	if int(flagInt.Int64()) == 0 {
		fmt.Println("encrypting with CBC")
		encryptedByte = challenge10.EncryptCBC(string(GetRandomBytes(16)),string(GetRandomBytes(16)),input)
	} else {
		fmt.Println("encrypting with EBC")
		encryptedByte = challenge10.EncryptEBC(string(GetRandomBytes(16)),input)
	}
	
	return 
}

