package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"main/src/engine"
)

func main() {
	var e engine.Engine

	e.Init()
	e.Load()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	e.Run()
	e.Unload()
	e.Close()
}
