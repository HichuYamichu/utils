package handler

import (
	"time"

	"github.com/hichuyamichu-me/utils/app/store"
	"github.com/jinzhu/gorm"
)

const lifetime = time.Second * 60 * 30

type Handler struct {
	FS *store.FS
	DB *gorm.DB
}

func New() *Handler {
	return &Handler{FS: &store.FS{}}
}
