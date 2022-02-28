package model

import "gorm.io/gorm"

// Book 书的属性
type Book struct {
	gorm.Model
	Title     string `json:"title"`
	PublishBy string `json:"publish_by"`
	Edition   string `json:"edition"`
	Author    string

	Value []Value `gorm:"ForeignKey:BookId;ASSOCIATION_FOREIGNKEY:ID"`
}

type Images struct {
	gorm.Model
	ValueId uint
	Large   string
	Medium  string
	Tiny    string
}

type Value struct {
	gorm.Model
	BookId string
	Price  float64  //价格
	Wear   int      //磨损层度
	Images []Images `gorm:"ForeignKey:ValueId;ASSOCIATION_FOREIGNKEY:ID"`
}

// Cart 购物车
type Cart struct {
	gorm.Model
	Uid       string `json:"uid"`
	BookId    string `gorm:"book_id" json:"book_id"`
	Wear      int    `json:"wear"`
	Transport string `json:"transport"`
}

type UserClickRecord struct {
	gorm.Model
	Uid    string
	BookId string
}
