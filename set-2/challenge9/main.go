package challenge9


import (

)

func main () {

}


func AddPadding(block string, blockSize int) ( string) {

	nPads := blockSize - len(block)
	for i := 0; i<nPads; i++ {
		block = block + "\x04"
	}
	
	return block
}

