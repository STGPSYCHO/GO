package models

type Product struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Picture string `json:"pic_link"`
}
