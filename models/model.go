package models

type Model struct {
	Id int64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
}
