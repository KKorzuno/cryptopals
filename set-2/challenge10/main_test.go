package main


import "testing"

type TestDataItem struct {
	block []byte
	blockSize int
	expectedOutputBlock []byte
}



func TestAddPadding( t *testing.T ) {

	// input-result data items
	dataItems := []TestDataItem{
		{ []byte("YELLOW SUBMARINE"),20,[]byte("YELLOW SUBMARINE\x04\x04\x04\x04")},
		{ []byte("YELLOW SUBMARINE"),15,[]byte("YELLOW SUBMARINE\x04\x04\x04\x04\x04\x04\x04\x04\x04\x04\x04\x04\x04\x04")},
		{ []byte("YELLOW SUBMARINE"),16,[]byte("YELLOW SUBMARINE")},
		{ []byte("YELLOW SUBMARINE"),4,[]byte("YELLOW SUBMARINE")},
		{ []byte("YELLOW SUBMARINE"),5,[]byte("YELLOW SUBMARINE\x04\x04\x04\x04")},



	}

	for _, item := range dataItems {
		result := AddPadding(item.block, item.blockSize)

		if len(result) != len(item.expectedOutputBlock) {
			t.Errorf( "addPadding() with args %v %v : FAILED, expected value '%v', but got '%v'",  (item.block), item.blockSize,  (item.expectedOutputBlock), result)
		} else {
			t.Logf( "addPadding() with args %v %v : PASSED, expected an error and got an error '%v'", (item.block), item.blockSize,  (item.expectedOutputBlock))
		}
		
	}
}

func TestXOROnBytes( t *testing.T ) {

	bytes1 := []byte("1234")
	bytes2 := []byte("1234")
	result := XOROnBytes(bytes1,bytes2)

	if len(result) != len(bytes1) {
		t.Errorf("Result not equal lenght as input input:%v result:%v",bytes1,result)
	}

}

func TestEncodeDecode( t *testing.T) {
	iv := "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	key:= "0000000000000000"
	bytes := []byte ("yellow submarineyellow submarine")

	encrypted := encryptCBC(iv,key,bytes)
	decrypted := decryptCBC(iv, key, encrypted)

	for i:=0;i<len(bytes);i++ {
		if bytes[i] != decrypted[i] {
			t.Error()
		}
	}

	t.Logf("bytes: %v",bytes)
	t.Logf("dec: %v",decrypted)
}

