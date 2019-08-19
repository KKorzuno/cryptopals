package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"errors"
)

func Hex_string_to_base64(msg string) (output string, err error) {
	decoded_hex_string, err := hex.DecodeString(msg)
	output = base64.StdEncoding.EncodeToString(decoded_hex_string)
	//fmt.Println(output)
	return
}

func Fixed_XOR_on_strings(hex_string, hex_string2 string) (output_hex_string string, err error) {
	decoded_hex_string, err := hex.DecodeString(hex_string)
	decoded_hex_string2, err := hex.DecodeString(hex_string2)

	//fmt.Printf("type: %T", decoded_hex_string)
	//base64_str, myerr := Hex_string_to_base64(mymsg)
	if len(decoded_hex_string) != len(decoded_hex_string2) {
		err = errors.New("Hex strings are not of equal lenght")		
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
	output_hex_string = hex.EncodeToString(output)	
	//fmt.Println(encoded_hex)
	return
}

func main(){
	xor_output,myerr := Fixed_XOR_on_strings("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if myerr != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(xor_output)
}
