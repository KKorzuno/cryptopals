package main

import (

	"fmt"
	"cryptopals/supportfunctions"
	"encoding/hex"
	"log"
)

func main(){
	hex1:= "1c0111001f010100061a024b53535009181c"
	hex2:= "686974207468652062756c6c277320657965"
	byte1, err1 := supportfunctions.HexStringToBytes(hex1)
	byte2, err2 := supportfunctions.HexStringToBytes(hex2)
	if err1!=nil {log.Fatal(err1)}
	if err2!=nil {log.Fatal(err2)}
	xorOutput := supportfunctions.XOROnBytes(byte1, byte2)
	fmt.Println(hex.EncodeToString(xorOutput))
}
