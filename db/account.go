package db

import (
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/types"
	merkle "github.com/cbergoon/merkletree"
)

// FetchSiblings retuns the siblings of an account leaf till root
func FetchSiblings() {

}

func StoreMT(mt merkle.MerkleTree) error {
	session := MgoSession.Copy()
	defer session.Close()
	if err := session.GetCollection(common.DATABASE, common.ACCOUNT_COLLECTION).Insert(mt); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}

func CreateAndStoreMT(accounts []types.AccountLeaf) {
	types.CreateTree()
}
