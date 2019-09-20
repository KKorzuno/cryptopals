package main

import "testing"

type TestDataItem struct {
	unknownString []byte
	unknownKey string
}

func TestDiscoverBlocksize( t *testing.T ) {
	dataItems := []TestDataItem{
		{[]byte("Helllo helllo helllo"),"YELLOW SUBMARINE"},
		// {[]byte("Helllo helllo helllo"),"1234"},
		// {[]byte("Helllo helllo helllo"),"12345"},
		// {[]byte("Helllo helllo helllo"),"123456"},
		// {[]byte("Helllo helllo helllo"),"123"},
		// {[]byte("Helllo helllo helllo"),"123"},
	}

	for _, item := range dataItems {
		keysize := discoverBlocksize(item.unknownString,item.unknownKey)
		if (keysize != len(item.unknownKey)) {
			t.Errorf("Problem: %v",item.unknownKey)
		}
	}

}