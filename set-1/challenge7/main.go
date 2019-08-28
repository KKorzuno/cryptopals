package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Print("STARTING THE PROGRAM\n")
	key := "YELLOW SUBMARINE"
	file, err := os.Open("7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytesFromFileInBase64, err := ioutil.ReadAll(file)
	BytesFromFile, err := base64.StdEncoding.DecodeString(string(bytesFromFileInBase64))
	keysize := len(key)
	fmt.Println(keysize)
	fmt.Println(len(BytesFromFile))
	fmt.Println(len(BytesFromFile) / keysize)
	fmt.Println("****************************************************")

	mycipher, _ := aes.NewCipher([]byte(key))

	BytesFromFileIn2D := make([][]byte, len(BytesFromFile)/keysize)
	decodedBytesIn2D := make([][]byte, len(BytesFromFile)/keysize)
	fmt.Println(string(BytesFromFile))
	fmt.Println("***************************************************")

	for i := 0; i < len(BytesFromFileIn2D); i++ {
		BytesFromFileIn2D[i] = BytesFromFile[keysize*i : keysize*(i+1)]
		decodedBytesIn2D[i] = make([]byte, len(key))
		mycipher.Decrypt(decodedBytesIn2D[i], BytesFromFileIn2D[i])
		fmt.Print(string(decodedBytesIn2D[i]))
	}


}