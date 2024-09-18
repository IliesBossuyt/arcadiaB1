package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.Sprite = rl.LoadTexture("textures/entities/soldier/Soldier-Idle.png")
	e.Sprites["HEALTHBAR"] = rl.LoadTexture("textures/entities/soldier/healthbar.png")
	e.Sprites["BACKGROUND"] = rl.LoadTexture("textures/map/tilesets/ecran.png")
	e.Sprites["BACKGROUNDPAUSE"] = rl.LoadTexture("textures/map/tilesets/pause.png")
	e.Sprites["HISTORY"] = rl.LoadTexture("textures/map/tilesets/intro.png")
	e.Sprites["DEAD"] = rl.LoadTexture("textures/map/tilesets/dead.png")
}

func (e *Engine) Unload() {
	// On libère les textures chargées, le joueur, la map, les monstres, etc...
	rl.UnloadTexture(e.Player.Sprite)

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}
}
