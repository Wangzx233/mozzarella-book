package logic

// Book 书的属性
type Book struct {
	BookId    string `json:"book_id"`
	Title     string `json:"title"`
	PublishBy string `json:"publish_by"`
	Edition   string `json:"edition"`
	Author    string `json:"author"`

	Value []Value `json:"value"`
}

type Value struct {
	BookId string   `json:"book_id"`
	Price  string   `json:"price"` //价格
	Wear   int      `json:"wear"`  //磨损层度
	Images []Images `json:"images"`
}
type Images struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Tiny   string `json:"tiny"`
}

// Cart 购物车
type Cart struct {
	Wear      string `json:"wear"`
	Transport string `json:"transport"`
	Book
}
