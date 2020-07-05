package core

import (
	"encoding/hex"

	"github.com/BOPR/contracts/rollup"
)

type AccountMerkleProof struct {
	Account  UserAccount
	Siblings []UserAccount
}

func NewAccountMerkleProof(account UserAccount, siblings []UserAccount) AccountMerkleProof {
	return AccountMerkleProof{Account: account, Siblings: siblings}
}

func (m *AccountMerkleProof) ToABIVersion() (AccMP rollup.TypesAccountMerkleProof, err error) {
	// create siblings
	var siblingNodes [][32]byte
	for _, s := range m.Siblings {
		siblingNodes = append(siblingNodes, s.HashToByteArray())
	}

	account, err := m.Account.ToABIAccount()
	if err != nil {
		return
	}

	AccMP = rollup.TypesAccountMerkleProof{
		AccountIP: rollup.TypesAccountInclusionProof{
			PathToAccount: StringToBigInt(m.Account.Path),
			Account:       account,
		},
		Siblings: siblingNodes,
	}

	return AccMP, nil
}

type PDAMerkleProof struct {
	Path      string
	PublicKey string
	Siblings  []UserAccount
}

func NewPDAProof(path string, publicKey string, siblings []UserAccount) PDAMerkleProof {
	return PDAMerkleProof{PublicKey: publicKey, Siblings: siblings, Path: path}
}

func (m *PDAMerkleProof) ToABIVersion() rollup.TypesPDAMerkleProof {
	// create siblings
	var siblingNodes [][32]byte
	for _, s := range m.Siblings {
		siblingNodes = append(siblingNodes, s.PubkeyHashToByteArray())
	}
	pubkey, err := hex.DecodeString(m.PublicKey)
	if err != nil {
		panic(err)
	}
	return rollup.TypesPDAMerkleProof{
		Pda: rollup.TypesPDAInclusionProof{
			PathToPubkey: StringToBigInt(m.Path),
			PubkeyLeaf:   rollup.TypesPDALeaf{Pubkey: pubkey},
		},
		Siblings: siblingNodes,
	}
}
