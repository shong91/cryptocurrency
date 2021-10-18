package blockchain

import (
	"sync"
)


type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height 		 int `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string){
	// save on DB
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

// singleton pattern: share ONLY 1 INSTANCE in application
func Blockchain() *blockchain {
	if b == nil {
		// sync.Once: method which calls only once
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis")
		})
	}
	return b
}
