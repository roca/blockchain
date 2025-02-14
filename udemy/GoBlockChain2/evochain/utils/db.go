package utils

import (
	"evochain/blockchain"
	"evochain/constants"

	"github.com/syndtr/goleveldb/leveldb"
)

func PutIntoDb(bc blockchain.Blockchain) error {
	db, err := leveldb.OpenFile(constants.BLOCKCHAIN_DB_PATH, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}
