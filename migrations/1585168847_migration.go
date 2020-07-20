package migrations

import (
	types "github.com/BOPR/core"
	"github.com/jinzhu/gorm"
)

// NOTE: The order is really important here in both up and down migrations
func init() {
	m := &Migration{
		ID: "1585168847",
		Up: func(db *gorm.DB) error {
			if !db.HasTable(&types.Tx{}) {
				db.CreateTable(&types.Tx{})
			}
			if !db.HasTable(&types.Batch{}) {
				db.CreateTable(&types.Batch{})
			}
			if !db.HasTable(&types.SyncStatus{}) {
				db.CreateTable(&types.SyncStatus{})
			}
			if !db.HasTable(&types.Params{}) {
				db.CreateTable(&types.Params{})
			}
			if !db.HasTable(&types.Token{}) {
				db.CreateTable(&types.Token{})
			}
			if !db.HasTable(&types.UserAccount{}) {
				db.CreateTable(&types.UserAccount{})
			}
			return nil
		},
		Down: func(db *gorm.DB) error {
			db.DropTableIfExists(&types.Tx{})
			db.DropTableIfExists(&types.Batch{})
			db.DropTableIfExists(&types.Params{})
			db.DropTableIfExists(&types.SyncStatus{})
			db.DropTableIfExists(&types.Token{})
			db.DropTableIfExists(&types.UserAccount{})
			return nil
		},
	}

	// add migration to list
	addMigration(m)
}
