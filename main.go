package main

import (
	"log"
	"main/src/engine"
	"net/http"
	_ "net/http/pprof"
	"os"
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

func godmode() {
    var godmode bool

    // Vérifier si l'argument -godmode est passé
    for _, arg := range os.Args {
        if arg == "--godmode" {
            godmode = true
            break
        }
    }

    var e engine.Engine

    // Initialiser l'engin du jeu
    e.Init()

    // Activer le mode godmode si l'argument est présent
    if godmode {
        log.Println("Godmode activé")
        e.EnableGodMode()
    }

    // Charger les ressources et démarrer le jeu
    e.Load()
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    e.Run()
    e.Unload()
    e.Close()
}
