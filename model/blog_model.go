package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Judul   string `json:"judul" form:"judul"`
	Konten  string `json:"konten" form:"konten"`
	Id_user uint   `json:"id_user" form:"id_user"`
	User    User   `json:"user" gorm:"foreignKey:Id_user"`
}
