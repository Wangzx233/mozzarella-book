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
	ValueId uint     `json:"value_id"`
	BookId  string   `json:"book_id"`
	Price   float64  `json:"price"` //价格
	Wear    int      `json:"wear"`  //磨损层度
	Images  []Images `json:"images"`
}
type Images struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Tiny   string `json:"tiny"`
}

// Cart 购物车
type Cart struct {
	Wear      int    `json:"wear"`
	Transport string `json:"transport"`
	Book
}
