package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/shong91/cryptocurrency/db"
	"github.com/shong91/cryptocurrency/utils"
)


type Block struct {
	// hash = data + prevHash
	// data -> hash is one-way function
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) toBytes() []byte {
	// return bytes
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	err := encoder.Encode(b)
	utils.HandleErr(err)

	return blockBuffer.Bytes()
} 

func (b *Block) persist(){
	db.SaveBlock(b.Hash, b.toBytes())
}

func createBlock(data string, prevHash string, height int) *Block{
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	
	// create hash and block hash
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}