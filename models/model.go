package models

type Model struct {
	Id uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
}
