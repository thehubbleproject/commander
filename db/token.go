package db

import (
	"github.com/BOPR/types"
)

func (db *DB) AddToken(t types.Token) error {
	return db.Instance.Create(t).Error
}

func (db *DB) GetTokenByID(id uint) (token types.Token, err error) {
	err = db.Instance.Where("token_id = ?", id).First(&token).Error
	if err != nil {
		return
	}
	return
}
