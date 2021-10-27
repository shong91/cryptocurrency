package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/shong91/cryptocurrency/db"
	"github.com/shong91/cryptocurrency/utils"
)

const difficulty int = 2

type Block struct {
	// hash = data + prevHash
	// data -> hash is one-way function
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
	Difficulty int `json:"difficulty"`
	Nonce	int `json:"nonce"`
}

func (b *Block) persist(){
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte){
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error){
	// get blockbytes
	blockbytes := db.Block(hash)
	if blockbytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockbytes)
	return block, nil

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