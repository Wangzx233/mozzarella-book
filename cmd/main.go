package main

import (
	"mozzarella-book/model"
	"mozzarella-book/route"
	"mozzarella-book/rpc"
)

func main() {
	model.InitMysql()

	rpc.InitRegister()
	route.InitRoute()
}
