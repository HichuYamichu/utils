package store

import "github.com/jinzhu/gorm"

type Store struct {
	FS *fs
	DB *gorm.DB
}

func New() *Store {
	return &Store{FS: &fs{}}
}
