package main

import (
	"cryptopals/supportfunctions"
	"fmt"
	"log"
)



func main(){
	mymsg := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64str, myerr := supportfunctions.HexStringToBase64(mymsg)
	if myerr != nil {
		log.Fatal("problems when encoding from hex to base64")
		return
	}
	fmt.Println(base64str)
	return
}
