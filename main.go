package main

import (
	_ "BscDapp/boot"
	_ "BscDapp/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
