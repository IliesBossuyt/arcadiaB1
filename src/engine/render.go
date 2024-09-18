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

func (e *Engine) HistoryRendering() {
	rl.DrawTexture(e.Sprites["HISTORY"], 0, 0, rl.RayWhite)
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()
	e.Displaydealer()
	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Money : %d $", int32(e.Player.Money)), int32(rl.GetScreenWidth())-rl.MeasureText("Money :", 515), int32(rl.GetScreenHeight())/50, 30, rl.RayWhite)     // Print the money
	rl.DrawText(fmt.Sprintf("Health : %d", int32(e.Player.Health)), int32(rl.GetScreenWidth())-rl.MeasureText("Health :", 500), int32(rl.GetScreenHeight())/18, 30, rl.RayWhite)    // Print the health
	rl.DrawText(fmt.Sprintf("Stamina : %d", int32(e.Player.Stamina)), int32(rl.GetScreenWidth())-rl.MeasureText("Stamina :", 443), int32(rl.GetScreenHeight())/11, 30, rl.RayWhite) // Print the stamina

}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(e.Sprites["BACKGROUNDPAUSE"], 0, 0, rl.RayWhite)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 20, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Q]/[A] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

}
func (e *Engine) InvRendering() {
	rl.ClearBackground(rl.Gray)
	if len(e.Dealer.Inv) == 0 {
		rl.DrawText("Votre inventaire est vide.", 100, 150, 20, rl.RayWhite)

	} else {
		for i, item := range e.Dealer.Inv {
			text := item.Name
			rl.DrawText(text, 100, int32(150+i*170), 100, rl.RayWhite)
			rl.DrawTexturePro(
				item.Sprite,
				rl.NewRectangle(0, 0, 64, 64),
				rl.NewRectangle(100, float32(200+i*150), 125, 125),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White)
		}
		rl.DrawText(fmt.Sprintf("Argent : %d", e.Player.Money), 600, 100, 20, rl.RayWhite)
	}
}
func (e *Engine) RenderItems() {
	if !e.Player.Alive {
		return // Ne pas afficher les items si le joueur est mort
	}

	// Parcourir l'inventaire du joueur et afficher les items en haut à gauche
	for i, item := range e.Player.Inventory {
		itemText := fmt.Sprintf("%s", item.Name)
		rl.DrawText(itemText, 10, int32(100+i*100), 20, rl.RayWhite)
	}
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

func (e *Engine) RenderDialogDealer(d entity.Dealer, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(d.Position.X),
		int32(d.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) Normalexplanation(m entity.Dealer, sentence string) {
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
func (e *Engine) Displaydealer() {
	rl.DrawTexturePro(
		e.Dealer.Sprite, //normal
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Dealer.Position.X, e.Dealer.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

    for _, item := range e.Dealer.Inv {
        text := item.Name
        rl.DrawText(text, 100, int32(150+50), 20, rl.RayWhite)
    
       
       }
}
