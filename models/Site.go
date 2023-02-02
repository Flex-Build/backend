package models

type Site struct {
	Name string `json:"name" gorm:"primaryKey"`
}
