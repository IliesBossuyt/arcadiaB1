package engine

import (
	"fmt"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/GTA San Andreas Theme Song Full ! !.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
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
		e.Music = rl.LoadMusicStream("sounds/music/GTA San Andreas Theme Song Full ! !.mp3")
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
	if rl.IsKeyPressed(rl.KeyTab){
		e.StateEngine = INGAME
	
	}
}

	


func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()

}

func (e *Engine) MonsterCollisions() {

	for _, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-20 &&
			monster.Position.X < e.Player.Position.X+20 &&
			monster.Position.Y > e.Player.Position.Y-20 &&
			monster.Position.Y < e.Player.Position.Y+20 {

			if monster.Name == "claude" {
				e.NormalTalk(monster, "Bonjour")
				if rl.IsKeyPressed(rl.KeyE) {
					//lancer un combat ?
				}
			}
		} else {
			//...
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
			if rl.IsKeyPressed(rl.KeyL){				
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
