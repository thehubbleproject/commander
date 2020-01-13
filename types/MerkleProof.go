package types

import (
	"github.com/BOPR/contracts/rollup"
)

type MerkleProof struct {
	Account  AccountLeaf   `bson:"account"`
	Siblings []AccountLeaf `bson:"siblings"`
}

func NewMerkleProof(account AccountLeaf, siblings []AccountLeaf) MerkleProof {
	return MerkleProof{Account: account, Siblings: siblings}
}

func (m *MerkleProof) ToABIVersion() rollup.DataTypesMerkleProof {
	return rollup.DataTypesMerkleProof{
		Account:  m.Account.ToABIAccount(),
		Siblings: AccsToLeafHashes(m.Siblings),
	}
}
