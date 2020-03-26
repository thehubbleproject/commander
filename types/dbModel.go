package types

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// DBModel represents database model
//
//
// Example:
// type User struct {
// 	DBModel
// }
//
// func (user *User) BeforeCreate(scope *gorm.Scope) error {
// 	if err := user.DBModel.BeforeCreate(scope); err != nil {
// 		return err
// 	}
// 	return nil
// }
type DBModel struct {
	ID        string     `json:"-" gorm:"primary_key;size:100;default:'6ba7b810-9dad-11d1-80b4-00c04fd430c8'"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// BeforeCreate sets id
func (model *DBModel) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("id", uuid.NewV4().String())
	if err != nil {
		return err
	}
	return nil
}
