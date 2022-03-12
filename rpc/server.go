package rpc

import (
	"context"
	"mozzarella-book/model"
	rpc "mozzarella-book/rpc/pb"
)

type MozzarellaBook struct {
}

func (s *MozzarellaBook) Ping(context.Context, *rpc.PingRequest) (*rpc.PingReply, error) {
	return &rpc.PingReply{Pong: "Pong"}, nil
}

func (s *MozzarellaBook) GetBookInfoByBookId(ctx context.Context, req *rpc.BookId) (*rpc.BookInfo, error) {
	//通过model层函数获取书籍
	book, err := model.SearchBookWithBookId(req.BookId)

	//转换value类型
	var BookInfo = rpc.BookInfo{
		Title:     book.Title,
		PublishBy: book.PublishBy,
		Edition:   book.Edition,
	}
	for _, value := range book.Value {
		var rValue rpc.Value
		rValue.Wear = int64(value.Wear)
		rValue.Price = float32(value.Price)
		//转换image
		for _, image := range value.Images {
			var rImage = rpc.Image{
				Large:  image.Large,
				Medium: image.Medium,
				Tiny:   image.Tiny,
			}
			rValue.Images = append(rValue.Images, &rImage)
		}

		BookInfo.Values = append(BookInfo.Values, &rValue)
	}

	return &BookInfo, err
}
