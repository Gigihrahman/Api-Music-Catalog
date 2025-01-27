package trackactivities

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepositoy(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}
