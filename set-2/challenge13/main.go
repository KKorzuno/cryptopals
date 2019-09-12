package main

import (
	"cryptopals/set-2/challenge10"
	//"encoding/base64"
	"cryptopals/set-1/challenge8"
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
	fmt.Println("ABCD")
	estimatedKeyLength := DiscoverBlocksize()
	fmt.Println(estimatedKeyLength)
	//modifiedEncryptedBytes := nastyAttacker(encryptedBytes)

	//systemToFool(secretKey, modifiedEncryptedBytes)

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
	decryptedBytes := challenge10.DecryptEBC(secretKey, encryptedBytes)
	receviedProfile := loginCredsParser(string(decryptedBytes))
	fmt.Println("Im the system that does not want to be fooled and I received the following profile:")
	fmt.Println(receviedProfile)
}

func nastyAttacker(encryptedBytes []byte) (modifiedencryptedBytes []byte) {
	//challenge8.
	return
}

func encrypter(email string) []byte {

	newProfile := profileFor(email)
	return challenge10.EncryptEBC(secretKey, []byte(newProfile))
}

func DiscoverBlocksize() (guessedBlockSize int) {
	listOfAs := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	for i := 0; i < 15; i++ {
		listOfAs = strings.Repeat("A", 16*2)
		fmt.Println("IN DISCOVER BLOCK SIZE")
		minDistance := challenge8.GetMinDistanceInKeysizeMultiComparison(encrypter(listOfAs), 16, 2)

		if minDistance == 0 {
			guessedBlockSize = 16
			return
		}
	}
	fmt.Println("AFTER 16 BYTE CHECK")
	listOfAsForKeySize24 := strings.Repeat("A", 24*3)
	for i := 0; i < 23; i++ {
		listOfAs = listOfAs + "A"
		if challenge8.CheckIfMinDistanceIsEqual3Times(encrypter(listOfAsForKeySize24), 24) {
			guessedBlockSize = 24
			return
		}
	}

	fmt.Println("AFTER 24 BYTE CHECK")
	listOfAsForKeySize32 := strings.Repeat("A", 32*3)
	for i := 0; i < 31; i++ {
		listOfAs = listOfAs + "A"
		if challenge8.CheckIfMinDistanceIsEqual3Times(encrypter(listOfAsForKeySize32), 32) {
			return 32
		}
	}
	fmt.Println("AFTER 32 BYTE CHECK")
	return -1
}
