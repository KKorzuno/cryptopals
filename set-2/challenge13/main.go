package main

import (
	"cryptopals/supportfunctions"

	//"encoding/base64"
	"fmt"
	//"cryptopals/set-2/challenge12"
	"strconv"
	"strings"
)

type loginCreds struct {
	email string
	uid   int
	role  string
}

var secretKey string

func main() {
	secretKey = "YELLOW SUBMARINE"

	encryptedBytes := encrypter("krzysztof@masterofdistaster.com")

	systemToFool(secretKey, encryptedBytes)
	PrintSliceNicely(encryptedBytes, len(secretKey))
	estimatedKeyLength, neededPadding, doubleBytePosition := supportfunctions.DiscoverBlocksize()
	fmt.Printf("Estimated key lenght: %v, Needed Padding: %v, Double byte found on position (counting from 0): %v\n",estimatedKeyLength, neededPadding, doubleBytePosition)
	modifiedEncryptedBytes := nastyAttacker(encryptedBytes, estimatedKeyLength, neededPadding, doubleBytePosition)
	PrintSliceNicely(modifiedEncryptedBytes, len(secretKey))
	//systemToFool(secretKey, modifiedEncryptedBytes)
	systemToFool(secretKey, modifiedEncryptedBytes)
}

func profileFor(email string) (encodedLogin string) {

	var newLoginCreds loginCreds

	newLoginCreds.email = escapeInput(email)
	newLoginCreds.uid = 10
	newLoginCreds.role = "user"

	return "email=" + newLoginCreds.email + "&uid=" + strconv.Itoa(newLoginCreds.uid) + "&role=" + newLoginCreds.role
}

func escapeInput(stringToParse string) (escapedString string) {
	escapedString = strings.ReplaceAll(stringToParse, "&", "")
	escapedString = strings.ReplaceAll(escapedString, "=", "")
	return
}

func loginCredsParser(concatenatedCreds string) (receivedLoginCreds loginCreds) {
	fmt.Println(concatenatedCreds)
	credParts := strings.Split(concatenatedCreds, "&")
	if len(credParts) < 3 {
		fmt.Println("wrong input for parser")
	}
	tempText := make([][]string, len(credParts))
	for i, v := range credParts {
		tempText[i] = strings.Split(v, "=")

	}
	receivedLoginCreds.email = tempText[0][1]
	receivedLoginCreds.uid, _ = strconv.Atoi(tempText[1][1])
	receivedLoginCreds.role = tempText[2][1]
	return
}

func systemToFool(secretKey string, encryptedBytes []byte) {
	//decryptedBytes := challenge10.DecryptEBC(secretKey, encryptedBytes)
	decryptedBytes := supportfunctions.Decrypt(secretKey, encryptedBytes)
	receviedProfile := loginCredsParser(string(decryptedBytes))
	fmt.Println("Im the system that does not want to be fooled and I received the following profile:")
	fmt.Println(receviedProfile)
}

func nastyAttacker(encryptedBytes []byte, estimatedKeyLength int, neededPadding int, doubleBytePosition int) (modifiedEncryptedBytes []byte) {
	adminPadded := "admin" + strings.Repeat("\x04",estimatedKeyLength-len("admin"))
	hostileAdminProfileEmail := strings.Repeat("A", neededPadding) + adminPadded
	fmt.Println(hostileAdminProfileEmail)
	tempBytesIn2D := supportfunctions.SliceBytesInto2D(encrypter(hostileAdminProfileEmail),estimatedKeyLength)
	
	//18+email%estimatedKeyLength=0  -- 14 bytes long mod 16 
	emailForPadding:="qwe@gmail.com"
	//modifiedEncryptedBytes
	bytesToSwapRole:= encrypter(emailForPadding)
	bytesToSwapRoleIn2D := sliceBytesInto2D(bytesToSwapRole, estimatedKeyLength)
	bytesToSwapRoleIn2D[2]=tempBytesIn2D[doubleBytePosition]
	PrintSliceNicely(bytesToSwapRole, estimatedKeyLength)
	modifiedEncryptedBytes=challenge10.FlattenPadded2DArray(bytesToSwapRoleIn2D)
	return
}

func encrypter(email string) []byte {

	newProfile := profileFor(email)
	fmt.Println(newProfile)
	return challenge10.EncryptEBC(secretKey, []byte(newProfile))
}


//PrintSliceNicely Transforms the slice into a 2D slice of lenght lenghtOfaRow, then prints it nicely
func PrintSliceNicely (inputByte []byte, lenghtOfaRow int) (){
	slicesToPrint := sliceBytesInto2D(inputByte, lenghtOfaRow)
	for _,v := range slicesToPrint { 
		fmt.Println(v)
	}
}
