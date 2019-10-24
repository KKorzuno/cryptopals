package main

import (
	"cryptopals/supportfunctions"
	//"cryptopals/set-1/challenge2"
	"encoding/hex"
	"fmt"
	"strings"
)
//Burning 'em, if you ain't quick and nimble
//I go crazy when I hear a cymbal

func main() {
	myString := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	result := applyRepeatingKeyXOR(myString,"ICE")

	fmt.Println(hex.EncodeToString(result))
	fmt.Println("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
}

func applyRepeatingKeyXOR(str string, repeatingKey string) ([]byte) {
	XORPAttern := strings.Repeat(repeatingKey,len(str)/len(repeatingKey))
	numberOfLeftovers := len(str)%len(repeatingKey)
	if( numberOfLeftovers != 0) {
		XORPAttern = XORPAttern + repeatingKey[0:numberOfLeftovers]
	}
	return supportfunctions.XOROnBytes([]byte(str), []byte(XORPAttern))
}