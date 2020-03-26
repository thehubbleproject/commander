package migrations

import (
	"github.com/BOPR/types"
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
			if !db.HasTable(&types.BatchModel{}) {
				db.CreateTable(&types.BatchModel{})
			}
			if !db.HasTable(&types.ListenerLog{}) {
				db.CreateTable(&types.ListenerLog{})
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
			db.DropTableIfExists(&types.BatchModel{})
			db.DropTableIfExists(&types.Params{})
			db.DropTableIfExists(&types.ListenerLog{})
			db.DropTableIfExists(&types.Token{})
			db.DropTableIfExists(&types.UserAccount{})
			return nil
		},
	}

	// add migration to list
	addMigration(m)
}
