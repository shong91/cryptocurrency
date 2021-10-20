package db

import (
	"github.com/boltdb/bolt"
	"github.com/shong91/cryptocurrency/utils"
)

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
	checkpoint = "checkpoint"
)

var db *bolt.DB

func DB() *bolt.DB{
	// init DB
	if db == nil {
		dbPointer, err := bolt.Open("blockchain.db", 0600, nil)
		utils.HandleErr(err)
		db = dbPointer

		// create bucket (=table)
		err = db.Update(func(t *bolt.Tx) error {
			// data / block bucket
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)

	}
	return db
}

func SaveBlock(hash string, data []byte){
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		// key, value = hash, data
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)

}

func SaveBlockchain(data []byte){
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		// key, value = hash, data
		err := bucket.Put([]byte(checkpoint), data) // save only newesthash and height
		return err
	})
	utils.HandleErr(err)
	
}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}