package main

type Oracle interface {
	prepareInputString() string
	encrypt() []byte
}



func NewOracle(unknownString []byte, unknownKey string) Oracle {
return Oracle{
	unknownString: unknownString,
	unknownKey: unknownKey
	}
}
