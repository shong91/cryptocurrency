package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

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

func (b *Block) mine(){
	target := strings.Repeat("0", b.Difficulty)
	for {
		// hash all blocks 
		blockAsString := fmt.Sprint(b)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString)))
		fmt.Printf("blockAsString:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", blockAsString, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			// find hash
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}

	}
}
 
func createBlock(data string, prevHash string, height int) *Block{
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
		Difficulty: difficulty,
		Nonce: 0,
	}
	block.mine()
	block.persist()
	return block
}