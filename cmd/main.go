package main

import (
	"mozzarella-book/model"
	"mozzarella-book/register"
	"mozzarella-book/route"
)

func main() {
	model.Init()
	register.InitRegister()
	route.InitRoute()
}
