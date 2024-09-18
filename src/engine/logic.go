package engine

import (
	"fmt"
	"main/src/entity"
	"time"

	//"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// 0 si l'histoire n'a pas encore été lu, 1 si l'histoire a été lu
var ReadHistory = 0

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/fairy-lands-fantasy-music-in-a-magical-forest-fantasy.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) && ReadHistory == 1 {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)
	}
	if rl.IsKeyPressed(rl.KeyEnter) && ReadHistory == 0 {
		e.StateMenu = HISTORY
		ReadHistory = 1
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

// Explication de l'histoire
func (e *Engine) HistoryLogic() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

var Stamina = false

func (e *Engine) InGameLogic() {

	// Dealer logic
	e.dealerCollisions()

	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
		e.Player.IsAlive = false
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
		e.Player.IsAlive = false
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
		e.Player.IsAlive = false
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
		e.Player.IsAlive = true
	}

	// Mouvement Shift
	if rl.IsKeyDown(rl.KeyW) && rl.IsKeyDown(rl.KeyLeftShift) {
		if e.Player.Stamina > 0 {
			e.Player.Position.Y -= e.Player.Speed + 1
			e.Player.Stamina -= 1
		}
	} else if e.Player.Stamina < 100 && !Stamina {
		Stamina = true
		go func() {
			for e.Player.Stamina < 100 {
				e.Player.Stamina += 1
				time.Sleep(500 * time.Millisecond)
			}
			Stamina = false
		}()
	}
	if rl.IsKeyDown(rl.KeyS) && rl.IsKeyDown(rl.KeyLeftShift) {
		if e.Player.Stamina > 0 {
			e.Player.Position.Y += e.Player.Speed * 2
			e.Player.Stamina -= 1
		}
	}
	if rl.IsKeyDown(rl.KeyA) && rl.IsKeyDown(rl.KeyLeftShift) {
		if e.Player.Stamina > 0 {
			e.Player.Position.X -= e.Player.Speed * 2
			e.Player.IsAlive = false
			e.Player.Stamina -= 1
		}
	}
	if rl.IsKeyDown(rl.KeyD) && rl.IsKeyDown(rl.KeyLeftShift) {
		if e.Player.Stamina > 0 {
			e.Player.Position.X += e.Player.Speed * 2
			e.Player.IsAlive = true
			e.Player.Stamina -= 1
		}
	}

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

// Inv check
func (e *Engine) InvLogic() {

	if rl.IsKeyPressed(rl.KeyI) {
		if len(e.Player.Inventory) > 0 {
			item := e.Player.Inventory[0]
			fmt.Printf("Vous avez utilisé %s.\n", item.Name)
			e.Player.Inventory = e.Player.Inventory[1:]
		} else {
			fmt.Println("Je n'ai pas d'items a échanger")
		}

	}
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateEngine = INGAME

	}
}

func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()

}

func (e *Engine) MonsterCollisions() {

	for i, monster := range e.Monsters {

		if monster.Position.X > e.Player.Position.X-100 &&
			monster.Position.X < e.Player.Position.X+100 &&
			monster.Position.Y > e.Player.Position.Y-100 &&
			monster.Position.Y < e.Player.Position.Y+100 {

			e.NormalTalk(monster, fmt.Sprintf("%d", monster.Health))
			if monster.Name == "claude" {
				if e.Player.Position.X < e.Monsters[i].Position.X+30 {
					e.Monsters[i].Position.X -= 3
				}
				if e.Player.Position.Y < e.Monsters[i].Position.Y+30 {
					e.Monsters[i].Position.Y -= 3
				}
				if e.Player.Position.Y > e.Monsters[i].Position.Y-30 {
					e.Monsters[i].Position.Y += 3
				}
				if e.Player.Position.X > e.Monsters[i].Position.X-30 {
					e.Monsters[i].Position.X += 3
				}

				if monster.Position.X > e.Player.Position.X-31 &&
					monster.Position.X < e.Player.Position.X+31 &&
					monster.Position.Y > e.Player.Position.Y-31 &&
					monster.Position.Y < e.Player.Position.Y+31 {
					if rl.IsKeyPressed(rl.KeyE) && e.Monsters[i].Health > 0 {
						e.Monsters[i].Health -= 10
						rl.PlaySound(e.Player.SwordSound)
						if e.Player.Position.X > e.Monsters[i].Position.X {
							e.Monsters[i].Position.X -= 30
						}
						if e.Player.Position.X < e.Monsters[i].Position.X {
							e.Monsters[i].Position.X += 30
						}
					}
				}

			}
		} else {
		}
	}
}
func (e *Engine) dealerCollisions() {

	if e.Dealer.Position.X > e.Player.Position.X-20 &&
		e.Dealer.Position.X < e.Player.Position.X+20 &&
		e.Dealer.Position.Y > e.Player.Position.Y-20 &&
		e.Dealer.Position.Y < e.Player.Position.Y+20 {

		if e.Dealer.Name == "yannis" {
			e.RenderDialogDealer(e.Dealer, "Bonjour")
			if rl.IsKeyPressed(rl.KeyL) {
				e.updatedealer()
				e.StateEngine = INV
			}
		}
	}
}
func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}
func (e *Engine) dealertalk(m entity.Dealer, sentence string) {
	e.Normalexplanation(m, sentence)
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyA) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) updatedealer() {
	if rl.IsKeyPressed(rl.KeyOne) {
		e.buyItem(0)
	}
	if rl.IsKeyPressed(rl.KeyTwo) {
		e.buyItem(1)
	}
	if rl.IsKeyPressed(rl.KeyThree) {
		e.buyItem(2)
	}

}

func (e *Engine) buyItem(index int) {
	item := e.Dealer.Inv[index]
	if index < 0 || index >= len(e.Dealer.Inv) {
		fmt.Println("Index invalide")
		return
	}

	if e.Player.Money >= item.Price {
		e.Player.Money -= item.Price
		e.Player.Inventory = append(e.Player.Inventory, item)
		fmt.Printf("Vous avez acheté %s pour %d pièces\n", item.Name, item.Price)
	} else {
		fmt.Println("Pas assez d'argent !")
	}
}
