package main

import (
	"mozzarella-book/model"
	"mozzarella-book/register"
	"mozzarella-book/route"
)

func main() {
	model.InitMysql()
	
	register.InitRegister()
	route.InitRoute()
}
