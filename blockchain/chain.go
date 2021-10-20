package blockchain

import (
	"fmt"
	"sync"

	"github.com/shong91/cryptocurrency/db"
	"github.com/shong91/cryptocurrency/utils"
)


type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height 		 int `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist(){
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string){
	// save on DB
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// singleton pattern: share ONLY 1 INSTANCE in application
func Blockchain() *blockchain {
	if b == nil {
		// sync.Once: method which calls only once
		once.Do(func() {
			b = &blockchain{"", 0}
			// search for checkpoint on the DB
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				// restore b from bytes 
				b.restore(checkpoint)
			}			
		})
	}
	fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
 	return b
}
 