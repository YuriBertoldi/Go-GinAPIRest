package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Id       int    `json:"ID"`
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
}
