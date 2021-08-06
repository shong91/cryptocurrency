package main

import (
	"fmt"

	"github.com/shong91/cryptocurrency/blockchain"
)


func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second")
	chain.AddBlock("Third")
	chain.AddBlock("Fourth")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}
