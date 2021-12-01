package main

import (
	_ "jsForward/boot"
	_ "jsForward/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}