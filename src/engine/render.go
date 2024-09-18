package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"

	"fmt"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}

func (e *Engine) HomeRendering() {
	rl.DrawTexture(e.Sprites["BACKGROUND"], 0, 0, rl.RayWhite)

}

func (e *Engine) HistoryRendering() {
	rl.DrawTexture(e.Sprites["HISTORY"], 0, 0, rl.RayWhite)
	rl.DrawText("[Enter] to Continue", int32(rl.GetScreenWidth())-rl.MeasureText("[Enter] to Continue", 32), int32(rl.GetScreenHeight())/1-35, 30, rl.RayWhite)
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText(fmt.Sprintf("Money : %d $", int32(e.Player.Money)), int32(rl.GetScreenWidth())-rl.MeasureText("Money :", 515), int32(rl.GetScreenHeight())/50, 30, rl.RayWhite)     // Print the money
	rl.DrawText(fmt.Sprintf("Health : %d", int32(e.Player.Health)), int32(rl.GetScreenWidth())-rl.MeasureText("Health :", 500), int32(rl.GetScreenHeight())/18, 30, rl.RayWhite)    // Print the health
	rl.DrawText(fmt.Sprintf("Stamina : %d", int32(e.Player.Stamina)), int32(rl.GetScreenWidth())-rl.MeasureText("Stamina :", 443), int32(rl.GetScreenHeight())/11, 30, rl.RayWhite) // Print the stamina

}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(e.Sprites["BACKGROUNDPAUSE"], 0, 0, rl.RayWhite)

}

func (e *Engine) GameOverRendering() {
	rl.DrawTexture(e.Sprites["DEAD"], 0, 0, rl.RayWhite)
	rl.DrawText("[Enter] to Respawn", int32(rl.GetScreenWidth())-rl.MeasureText("[Enter] to Respawn", 32), int32(rl.GetScreenHeight())/1-35, 30, rl.RayWhite)

}

func (e *Engine) RenderPlayer() {

	if e.Player.IsAlive {
		rl.DrawTexturePro(
			e.Player.Sprite, //normal
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)

	} else {
		rl.DrawTexturePro(
			e.Player.Sprite, // invertion horizontal
			rl.NewRectangle(0, 0, -100, 100),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}

}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		rl.DrawTexturePro(
			monster.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
