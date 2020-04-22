package core

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

func (m *MerkleProof) ToABIVersion(path int64) rollup.DataTypesAccountMerkleProof {
	return rollup.DataTypesAccountMerkleProof{
		AccountIP: m.Account.AccountInclusionProof(path),
		// Siblings:  AccsToLeafHashes(m.Siblings),
	}
}
