package model

import (
	"gorm.io/gorm"
)

type Credential struct {
	gorm.Model
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type Category struct {
	gorm.Model
	Name       string      `gorm:"type:varchar(255)"`
	Mahfudzots []Mahfudzot `gorm:"foreignKey:CategoryID"`
}

type Mahfudzot struct {
	gorm.Model
	Author     string `gorm:"type:varchar(255)"`
	Content    string `gorm:"type:text"`
	CategoryID uint
}
