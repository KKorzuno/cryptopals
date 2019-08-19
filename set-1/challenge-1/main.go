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

func main(){
	mymsg := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64_str, myerr := Hex_string_to_base64(mymsg)
	if myerr != nil {
		fmt.Println("error:", myerr)
		return
	}
	fmt.Println(base64_str)
	return
}
