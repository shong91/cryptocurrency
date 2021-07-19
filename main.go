package main

import "crypto/sha256"

type block struct {
	// hash = data + prevHash
	data     string
	hash     string
	prevHash string
}

func main() {
	// 초기 block
	genesisBlock := block{"Genesis block", "", ""}
	genesisBlock.hash = sha256.Sum256(genesisBlock.data + genesisBlock.hash)

}
