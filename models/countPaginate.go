package models

import "gorm.io/gorm"

type CountPaginate interface {
	Count(db *gorm.DB) int64
	Paginate(db *gorm.DB, limit int, offset int) interface{}
}
