package main

import (
	server "github.com/MasoudHeydari/golang-keep-note/controller"
)

func main() {
	server.NewServer().Run(":8000")
}
