package main

import (
	"MyGlog/model"
	"MyGlog/routers"
)

func main() {
	model.InitDb()

	routers.InitRouter()

}
