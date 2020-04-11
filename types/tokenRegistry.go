package types

type Token struct {
	DBModel
	// Token Address on eth chain.
	Address Address `json:"address"`

	// token ID allocated to the token
	TokenID uint64 `json:"tokenID"`
}
