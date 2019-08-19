package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Hex_string_to_base64(msg string) (output string, err error) {
	decoded_hex_string, err := hex.DecodeString(msg)
	output = base64.StdEncoding.EncodeToString(decoded_hex_string)
	//fmt.Println(output)
	return
}

func main() {
	hex_string := "1c0111001f010100061a024b53535009181c"
	hex_string2 := "686974207468652062756c6c277320657965"
	decoded_hex_string, err := hex.DecodeString(hex_string)
	decoded_hex_string2, err := hex.DecodeString(hex_string2)

	//fmt.Printf("type: %T", decoded_hex_string)
	//base64_str, myerr := Hex_string_to_base64(mymsg)
	if len(decoded_hex_string) != len(decoded_hex_string2) {
		fmt.Println("error")
		return
	}
	if err != nil {
		fmt.Println("error")
		return
	}
	//fmt.Printf("\n%v %v", decoded_hex_string, decoded_hex_string2)
	output:=make([]byte,len(decoded_hex_string))
	for i:=0; i<len(decoded_hex_string); i++ {
		output[i] = decoded_hex_string[i]^decoded_hex_string2[i]
	}
	//fmt.Printf("\n\n\nOUTPUT: %v\n", output)
	encoded_hex := hex.EncodeToString(output)	
	fmt.Println(encoded_hex)
	return
}
