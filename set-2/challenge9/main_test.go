package  challenge9 

import "testing"

type TestDataItem struct {
	block string
	blockSize int
	expectedOutputBlock string
}



func TestAddPadding( t *testing.T ) {

	// input-result data items
	dataItems := []TestDataItem{
		{ "YELLOW SUBMARINE",20,"YELLOW SUBMARINE\x04\x04\x04\x04"},
		{ "YELLOW SUBMARINE",15,"YELLOW SUBMARINE"},
		{ "YELLOW SUBMARINE",16,"YELLOW SUBMARINE"},

	}

	for _, item := range dataItems {
		result := addPadding(item.block, item.blockSize)

		if result != item.expectedOutputBlock {
			t.Errorf( "addPadding() with args %v %v : FAILED, expected an error but got value '%v'", []byte (item.block), item.blockSize, []byte (item.expectedOutputBlock))
		} else {
			t.Logf( "addPadding() with args %v %v : PASSED, expected an error and got an error '%v'", []byte (item.block), item.blockSize, []byte (item.expectedOutputBlock))
		}
		
	}
}