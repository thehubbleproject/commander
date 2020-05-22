package core

import (
	"github.com/BOPR/contracts/coordinatorproxy"
)

type AccountMerkleProof struct {
	Account  UserAccount   `bson:"account"`
	Siblings []UserAccount `bson:"siblings"`
}

func NewAccountMerkleProof(account UserAccount, siblings []UserAccount) AccountMerkleProof {
	return AccountMerkleProof{Account: account, Siblings: siblings}
}

func (m *AccountMerkleProof) ToABIVersion() coordinatorproxy.TypesAccountMerkleProof {
	// create siblings
	var siblingNodes [][32]byte
	for _, s := range m.Siblings {
		siblingNodes = append(siblingNodes, s.HashToByteArray())
	}

	return coordinatorproxy.TypesAccountMerkleProof{
		AccountIP: coordinatorproxy.TypesAccountInclusionProof{
			PathToAccount: StringToBigInt(m.Account.Path),
			Account:       m.Account.ToABIAccount(),
		},
		Siblings: siblingNodes,
	}
}
