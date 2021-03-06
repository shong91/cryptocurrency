package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/shong91/cryptocurrency/db"
	"github.com/shong91/cryptocurrency/utils"
)

type Block struct {
	// hash = data + prevHash
	// data -> hash is one-way function
	
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
	Difficulty int `json:"difficulty"`
	Nonce	int `json:"nonce"`
	Timestamp int `json:"timestamp"`
	Transactions []*Tx `json:"transactions"`
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
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("\n\n\nTarget:%s\nHash:%s\nNonce:%d\n\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			// find hash
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}

	}
}
 
func createBlock(prevHash string, height int) *Block{
	block := &Block{
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
		Difficulty: Blockchain().difficulty(),
		Nonce: 0,
		Transactions: []*Tx{makeCoinbaseTx("hong")},
	}

	block.mine()
	block.persist()
	return block
}