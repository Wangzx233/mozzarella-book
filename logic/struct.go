package logic

// Book 书的属性
type Book struct {
	BookId    string   `json:"book_id"`
	Title     string   `json:"title"`
	PublishBy string   `json:"publish_by"`
	Edition   string   `json:"edition"`
	Images    []Images `json:"images"`
	Value     Value    `json:"value"`
}

type Images struct {
	Large  string
	Medium string
	Tiny   string
}

type Value struct {
	Price string //价格
	Wear  int    //磨损层度
}
