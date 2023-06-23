package repository

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrRaceUpdate = errors.New("race condition update detected")
)

type Lockable interface {
	SetVersion(version int)
	GetVersion() int
}

type VersionModel struct {
	Versioning int `gorm:"not null;default:1"`
}

func (vm *VersionModel) SetVersion(version int) {
	vm.Versioning = version
}

func (vm *VersionModel) GetVersion() int {
	return vm.Versioning
}

// Update with optimistic locking. It will update all fields including zero value fields but omitting some fields if specified
func UpdateWithLock(db *gorm.DB, model Lockable, omitFields ...string) error {
	currentVersion := model.GetVersion()
	newVersion := currentVersion + 1
	model.SetVersion(newVersion)

	tx := db.Select("*").Omit(omitFields...).Where("versioning = ?", currentVersion).Updates(model)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return ErrRaceUpdate
	}

	return nil
}
