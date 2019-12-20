package store

import "github.com/jinzhu/gorm"

type Store struct {
	FS *fs
	db *gorm.DB
}

func New() *Store {
	return &Store{FS: &fs{}}
}
