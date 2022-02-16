package model

import "gorm.io/gorm"

// Book 书的属性
type Book struct {
	gorm.Model
	Title     string `json:"title"`
	PublishBy string `json:"publish_by"`
	Edition   string `json:"edition"`
	Author    string

	Value []Value `json:"value"`
}

type Images struct {
	gorm.Model
	BookId string
	Large  string
	Medium string
	Tiny   string
}

type Value struct {
	gorm.Model
	BookId string
	Price  string   //价格
	Wear   int      //磨损层度
	Images []Images `gorm:"ForeignKey:BookId;ASSOCIATION_FOREIGNKEY:Id"`
}

// Cart 购物车
type Cart struct {
	gorm.Model
	Uid       string
	BookId    string `gorm:"book_id"`
	Wear      string
	Transport string
}
