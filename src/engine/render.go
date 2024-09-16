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

	rl.DrawText("Home Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[Enter] to Play", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Play", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()

	rl.EndMode2D() // On finit le rendu camera
	
	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Money : %d", int32(e.Player.Money)), int32(rl.GetScreenWidth())-rl.MeasureText("Money :", 515), int32(rl.GetScreenHeight())/50, 30, rl.RayWhite)    // Print the money
	rl.DrawText(fmt.Sprintf("Health : %d", int32(e.Player.Health)), int32(rl.GetScreenWidth())-rl.MeasureText("Health :", 500), int32(rl.GetScreenHeight())/18, 30, rl.RayWhite) // Print the health

}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(e.Sprites["BACKGROUNDPAUSE"], 0, 0, rl.RayWhite)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 20, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Q]/[A] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

}

func (e *Engine) RenderPlayer() {
	
	if e.Player.IsAlive{
		rl.DrawTexturePro(
			e.Player.Sprite,//normal
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)

	}else{
			rl.DrawTexturePro(
			e.Player.Sprite,// invertion horizontal
			rl.NewRectangle(0, 0, -100, 100),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
			rl.Vector2{X: 0 , Y: 0},
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
