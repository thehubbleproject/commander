package migrations

import (
	"github.com/BOPR/types"
	"github.com/jinzhu/gorm"
)

func init() {
	m := &Migration{
		ID: "1582859979",
		Up: func(db *gorm.DB) error {
			if !db.HasTable(&types.Tx{}) {
				db.CreateTable(&types.Tx{})
			}
			if !db.HasTable(&types.AccountLeaf{}) {
				db.CreateTable(&types.AccountLeaf{})
			}
			if !db.HasTable(&types.Batch{}) {
				db.CreateTable(&types.Batch{})
			}
			if !db.HasTable(&types.BatchInfo{}) {
				db.CreateTable(&types.BatchInfo{})
			}
			return nil
		},
		Down: func(db *gorm.DB) error {
			db.DropTableIfExists(&types.Tx{})
			db.DropTableIfExists(&types.Batch{})
			db.DropTableIfExists(&types.AccountLeaf{})
			return nil
		},
	}

	// add migration to list
	addMigration(m)
}
