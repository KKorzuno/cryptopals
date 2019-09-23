package main

import (
	"encoding/base64"
	"fmt"
	// "cryptopals/set-2/challenge12"
	//"cryptopals/set-1/challenge8"
	"cryptopals/supportfunctions"
)

func main() {
	unknownBytebase64 := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
	aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
	dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
	YnkK`
	unknownByte, _ := base64.StdEncoding.DecodeString(unknownBytebase64)
	//unknownByte = unknownByte[0:1]
	unknownKey := "YELLOW SUBMARINE"

	sillyOracle:= supportfunctions.NewPrefixInputPostfixECBOracle("", string(unknownByte),unknownKey)
	decyphered:= supportfunctions.PaddingOracleAttack(sillyOracle)

	fmt.Println(decyphered)
}


