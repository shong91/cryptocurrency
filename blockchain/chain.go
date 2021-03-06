package blockchain

import (
	"fmt"
	"sync"

	"github.com/shong91/cryptocurrency/db"
	"github.com/shong91/cryptocurrency/utils"
)

const (
	defaultDifficulty int = 2
	difficultyInterval int = 5
	blockInterval int = 2 
	allowedRange int = 2 
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height 		 int `json:"height"`
	CurrentDifficulty int `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist(){
	db.SaveCheckPoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(){
	// save on DB
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block,_ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			// break when reach to Genesis block 
			break
		}
	}
	return blocks
	
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval - 1]
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60) // 분 단위 계산 
	expectedTime := difficultyInterval * blockInterval
	// strict 
	// if actualTime < expectedTime {
	// 	return b.CurrentDifficulty +1
	// } else if actualTime > expectedTime {
	// 	return b.CurrentDifficulty -1
	// }
	// return b.CurrentDifficulty
	
	// rough range 
	if actualTime <= (expectedTime - allowedRange){
		return b.CurrentDifficulty +1
	} else if actualTime >= (expectedTime + allowedRange){
		return b.CurrentDifficulty -1
	}
	return b.CurrentDifficulty


}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % difficultyInterval == 0 {
		// recalculate the difficulty
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

// 모든 거래 출력량을 리턴. 
func (b *blockchain) txOuts() []*TxOut{
	var txOuts []*TxOut
	blocks := b.Blocks()
	for _, block := range blocks {
		for _, tx := range block.Transactions {
			// 각각의 블록 안에 있는 모든 거래들의 출력값을 하나의 슬라이스 (txOuts) 에 담는다. 
			txOuts = append(txOuts, tx.TxOuts...)
		}
	}
	return txOuts
}

// 주소에 따른 거래 출력량 리턴. 
func (b *blockchain) TxOutsByAddress(address string) []*TxOut {
	var ownedTxOuts []*TxOut
	txOuts := b.txOuts()
	for _, txOuts := range txOuts {
		if txOuts.Owner == address {
			ownedTxOuts = append(ownedTxOuts, txOuts)
		}
	}
	return ownedTxOuts

}

func (b *blockchain) BalanceByAddress(address string) int {
	txOuts := b.TxOutsByAddress(address)
	var amount int 
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

// singleton pattern: share ONLY 1 INSTANCE in application
func Blockchain() *blockchain {
	if b == nil {
		// sync.Once: method which calls only once
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}
			// search for checkpoint on the DB
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock()
			} else {
				// restore b from bytes 
				b.restore(checkpoint)
			}			
		})
	}
	fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
 	return b
}
 