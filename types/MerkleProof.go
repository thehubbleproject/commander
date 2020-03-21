package types

import (
	"github.com/BOPR/contracts/rollup"
)

type MerkleProof struct {
	Account  UserAccount   `bson:"account"`
	Siblings []UserAccount `bson:"siblings"`
}

func NewMerkleProof(account UserAccount, siblings []UserAccount) MerkleProof {
	return MerkleProof{Account: account, Siblings: siblings}
}

func (m *MerkleProof) ToABIVersion() rollup.DataTypesMerkleProof {
	return rollup.DataTypesMerkleProof{
		Account:  m.Account.ToABIAccount(),
		Siblings: AccsToLeafHashes(m.Siblings),
	}
}
