package core

import (
	"github.com/BOPR/contracts/coordinatorproxy"
)

type AccountMerkleProof struct {
	Account  UserAccount
	Siblings []UserAccount
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

type PDAMerkleProof struct {
	Path      string
	PublicKey string
	Siblings  []UserAccount
}

func NewPDAProof(path string, publicKey string, siblings []UserAccount) PDAMerkleProof {
	return PDAMerkleProof{PublicKey: publicKey, Siblings: siblings, Path: path}
}

func (m *PDAMerkleProof) ToABIVersion() coordinatorproxy.TypesPDAMerkleProof {
	// create siblings
	var siblingNodes [][32]byte
	for _, s := range m.Siblings {
		siblingNodes = append(siblingNodes, s.PubkeyHashToByteArray())
	}
	pubkey, _ := ABIEncodePubkey(m.PublicKey)
	return coordinatorproxy.TypesPDAMerkleProof{
		Pda: coordinatorproxy.TypesPDAInclusionProof{
			PathToPubkey: StringToBigInt(m.Path),
			PubkeyLeaf:   coordinatorproxy.TypesPDALeaf{Pubkey: pubkey},
		},
		Siblings: siblingNodes,
	}
}
