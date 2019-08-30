package challenge10

import (
	"crypto/aes"
	// "encoding/base64"
	"fmt"
	"log"
	"os"
	"encoding/hex"
	"encoding/base64"
	"io/ioutil"
	"cryptopals/set-1/challenge2"
	// "cryptopals/set-2/challenge9"

)


func main () {

	key := "YELLOW SUBMARINE"
	file, err := os.Open("10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// bytesFromFile, err := ioutil.ReadAll(file)
	bytesFromFileInBase64, err := ioutil.ReadAll(file)
	bytesFromFile, err := base64.StdEncoding.DecodeString(string(bytesFromFileInBase64))

	dec := DecryptCBC("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", key, bytesFromFile)
	fmt.Println(string(dec))
}


func EncryptEBC(key string, bytesToEncrypt []byte) []byte{
	keysize := len(key)
	mycipher, _ := aes.NewCipher([]byte(key))
	
	bytesToEncryptIn2D := make([][]byte, len(bytesToEncrypt)/keysize)
	encryptedBytesIn2D := make([][]byte, len(bytesToEncrypt)/keysize)

	for i := 0; i < len(bytesToEncryptIn2D); i++ {
		bytesToEncryptIn2D[i] = bytesToEncrypt[keysize*i : keysize*(i+1)]
		encryptedBytesIn2D[i] = make([]byte, len(key))
		mycipher.Encrypt(encryptedBytesIn2D[i], bytesToEncryptIn2D[i])
	}

	var encryptedBytesFlat []byte
	for _,e := range encryptedBytesIn2D {
		encryptedBytesFlat = append(encryptedBytesFlat, e...)
	}

	return encryptedBytesFlat
}

func DecryptEBC(key string, bytesToDecrypt []byte) []byte{
	keysize := len(key)
	mycipher, _ := aes.NewCipher([]byte(key))
	
	bytesToDecryptIn2D := make([][]byte, len(bytesToDecrypt)/keysize)
	decryptedBytesIn2D := make([][]byte, len(bytesToDecrypt)/keysize)

	for i := 0; i < len(bytesToDecryptIn2D); i++ {
		bytesToDecryptIn2D[i] = bytesToDecrypt[keysize*i : keysize*(i+1)]
		decryptedBytesIn2D[i] = make([]byte, len(key))
		mycipher.Decrypt(decryptedBytesIn2D[i], bytesToDecryptIn2D[i])
	}

	var decryptedBytesFlat []byte
	for _,e := range decryptedBytesIn2D {
		decryptedBytesFlat = append(decryptedBytesFlat, e...)
	}

	return decryptedBytesFlat
}

func EncryptCBC(iv string, key string, bytesToEncrypt []byte) []byte{
	keysize := len(key)
	paddedBytesToEncrypt := AddPadding(bytesToEncrypt, keysize)
	mycipher, _ := aes.NewCipher([]byte(key))
	supportVector := []byte (iv)
	bytesToEncryptIn2D := make([][]byte, len(paddedBytesToEncrypt)/keysize)
	encryptedBytesIn2D := make([][]byte, len(paddedBytesToEncrypt)/keysize)

	// bytesToEncryptIn2D[len(bytesToEncryptIn2D)-1] = []byte(challenge9.AddPadding(string(bytesToEncryptIn2D[len(bytesToEncryptIn2D)-1]),keysize))

	for i := 0; i < len(bytesToEncryptIn2D); i++ {
		bytesToEncryptIn2D[i] = paddedBytesToEncrypt[keysize*i : keysize*(i+1)]
		bytesToEncryptIn2D[i] = XOROnBytes(bytesToEncryptIn2D[i], supportVector)
		encryptedBytesIn2D[i] = make([]byte, len(key))
		// fmt.Println(encryptedBytesIn2D[i], bytesToEncryptIn2D[i])
		mycipher.Encrypt(encryptedBytesIn2D[i], bytesToEncryptIn2D[i])
		supportVector=encryptedBytesIn2D[i]
	}

	var encryptedBytesFlat []byte
	for _,e := range encryptedBytesIn2D {
		encryptedBytesFlat = append(encryptedBytesFlat, e...)
	}

	return encryptedBytesFlat
}

func XOROnBytes (by1 []byte , by2 []byte ) []byte {
	str,err :=  challenge2.Fixed_XOR_on_hex_strings(hex.EncodeToString(by1),hex.EncodeToString(by2))
		if (err != nil) {fmt.Println(err)}
	res,err :=  hex.DecodeString(str)
	if (err != nil) {fmt.Println(err)}
	return res
}

func AddPadding(  bytes []byte, keysize int) []byte{
	if len(bytes)%keysize == 0 {return bytes}
	nPads := keysize - len(bytes)%keysize
	for i := 0; i<nPads; i++ {
		bytes = append(bytes ,[]byte("\x04")...)
	}
	
	return bytes

}

func DecryptCBC(iv string, key string, bytesToDecrypt []byte) []byte{
	keysize := len(key)
	mycipher, _ := aes.NewCipher([]byte(key))
	supportVector := []byte (iv)
	// fmt.Println( (supportVector))
	// fmt.Println( len(iv))

	bytesToDecryptIn2D := make([][]byte, len(bytesToDecrypt)/keysize)
	decryptedBytesIn2D := make([][]byte, len(bytesToDecrypt)/keysize)

	// bytesToEncryptIn2D[len(bytesToEncryptIn2D)-1] = []byte(challenge9.AddPadding(string(bytesToEncryptIn2D[len(bytesToEncryptIn2D)-1]),keysize))

	for i := 0; i < len(bytesToDecryptIn2D); i++ {
		bytesToDecryptIn2D[i] = bytesToDecrypt[keysize*i : keysize*(i+1)]
		decryptedBytesIn2D[i] = make([]byte, len(key))
		mycipher.Decrypt(decryptedBytesIn2D[i], bytesToDecryptIn2D[i])
		// fmt.Println( len(decryptedBytesIn2D[i]), len(supportVector))
		decryptedBytesIn2D[i] = XOROnBytes(decryptedBytesIn2D[i], supportVector)
		supportVector=bytesToDecryptIn2D[i]
	}

	var decryptedBytesFlat []byte
	for _,e := range decryptedBytesIn2D {
		decryptedBytesFlat = append(decryptedBytesFlat, e...)
	}

	return decryptedBytesFlat
}