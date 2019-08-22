package challenge5

import (
	"../challenge2"
	"encoding/hex"
	"fmt"
	"strings"
)
//Burning 'em, if you ain't quick and nimble
  
//I go crazy when I hear a cymbal

func main() {
	myString := "Burning 'em, if you ain't quick and nimble"
	result, error := ApplyRepeatingKeyXOR(myString,"ICE")

	fmt.Println(result, error)

}

func ApplyRepeatingKeyXOR(str string, repeatingKey string) (string, error) {
	XORPAttern := strings.Repeat(repeatingKey,len(str)/len(repeatingKey))
	numberOfLeftovers := len(str)%len(repeatingKey)
	if( numberOfLeftovers != 0) {
		XORPAttern = XORPAttern + repeatingKey[0:numberOfLeftovers]
	}

	return challenge2.Fixed_XOR_on_hex_strings(hex.EncodeToString([]byte(str)), hex.EncodeToString([]byte(XORPAttern))) 
}