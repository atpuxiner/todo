package common

import (
	"errors"
	"gorm.io/gorm"
)

func HandleError(tx *gorm.DB, msg string) (err error){
	if tx.RowsAffected == 0{
		err = errors.New(msg)
	}else {
		err = tx.Error
	}
	return
}
