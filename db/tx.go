package db 

import (
	"fmt"
	"github.com/BOPR/common"
	"github.com/BOPR/types"
)
// Insert tx into the DB
func InsertTx(t *types.Tx) error {
	session := MgoSession.Copy()
	defer session.Close()
	if err := session.GetCollection(common.DATABASE, common.TRANSACTION_COLLECTION).Insert(t); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}