package main

import (
	"srf/controller"
	"srf/srf"
)

func main() {
	server := srf.NewServer("127.0.0.1", 8080, controller.MAPPER)
	e := server.Start()
	if e != nil {
		println(e.Error())
	}
}
