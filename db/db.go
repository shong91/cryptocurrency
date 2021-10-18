package db

import (
	"github.com/boltdb/bolt"
	"github.com/shong91/cryptocurrency/utils"
)

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
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