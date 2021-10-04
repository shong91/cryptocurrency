package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	// hash = data + prevHash
	// data -> hash is one-way function
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	// slice of pointer(*block)
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)

	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash

}

func createBlock(data string) *Block {
	// data, hash, prevhash
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// Uppercase means it will be exported
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))

}

// singleton pattern: share ONLY 1 INSTANCE in application
func GetBlockchain() *blockchain {
	if b == nil {
		// sync.Once: method which calls only once
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
