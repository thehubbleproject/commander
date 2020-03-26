package types

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model

	// Token Address on eth chain.
	Address Address `json:"address"`

	// token ID allocated to the token
	TokenID uint `json:"tokenID"`
}
