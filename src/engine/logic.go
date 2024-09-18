package engine

import (
	"fmt"
	"main/src/entity"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// 0 si l'histoire n'a pas encore été lu, 1 si l'histoire a été lu
var ReadHistory = 0

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/GTA San Andreas Theme Song Full ! !.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	boutonXplay := int32(450)
	boutonYplay := int32(830)
	boutonLargeurplay := int32(200)
	boutonHauteurplay := int32(150)
	//Menus
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && ReadHistory == 1 {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXplay), float32(boutonYplay), float32(boutonLargeurplay), float32(boutonHauteurplay))) {
			e.StateMenu = PLAY
			e.StateEngine = INGAME
			rl.StopMusicStream(e.Music)
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && ReadHistory == 0 {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXplay), float32(boutonYplay), float32(boutonLargeurplay), float32(boutonHauteurplay))) {
			e.StateMenu = HISTORY
			ReadHistory = 1
		}
	}
	boutonXexit := int32(1270)
	boutonYexit := int32(830)
	boutonLargeurexit := int32(200)
	boutonHauteurexit := int32(150)
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXexit), float32(boutonYexit), float32(boutonLargeurexit), float32(boutonHauteurexit))) {
			e.IsRunning = false
		}
	}
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	boutonXresume := int32(730)
	boutonYresume := int32(1000)
	boutonLargeurresume := int32(170)
	boutonHauteurresume := int32(60)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXresume), float32(boutonYresume), float32(boutonLargeurresume), float32(boutonHauteurresume))) {
			e.StateEngine = INGAME
		}
	}
	boutonXmenu := int32(1200)
	boutonYmenu := int32(1000)
	boutonLargeuremenu := int32(170)
	boutonHauteuremenu := int32(60)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXmenu), float32(boutonYmenu), float32(boutonLargeuremenu), float32(boutonHauteuremenu))) {
			e.StateMenu = HOME
			rl.StopMusicStream(e.Music)
		}
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
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
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
		e.Player.IsAlive = false
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
		e.Player.IsAlive = true
	}

	// Mouvement Dash
	if rl.IsKeyDown(rl.KeyW) && rl.IsKeyPressed(rl.KeyLeftShift) {
		if e.Player.Stamina >= 25 {
			e.Player.Position.Y -= e.Player.Speed + 30
			e.Player.Stamina -= 25
		}
	} else if e.Player.Stamina < 100 && !Stamina {
		Stamina = true
		go func() {
			for e.Player.Stamina < 100 {
				e.Player.Stamina += 1
				time.Sleep(70 * time.Millisecond)
			}
			Stamina = false
		}()
	}
	if rl.IsKeyDown(rl.KeyS) && rl.IsKeyPressed(rl.KeyLeftShift) {
		if e.Player.Stamina >= 25 {
			e.Player.Position.Y += e.Player.Speed + 30
			e.Player.Stamina -= 25
		}
	}
	if rl.IsKeyDown(rl.KeyA) && rl.IsKeyPressed(rl.KeyLeftShift) {
		if e.Player.Stamina >= 25 {
			e.Player.Position.X -= e.Player.Speed + 30
			e.Player.IsAlive = false
			e.Player.Stamina -= 25
		}
	}
	if rl.IsKeyDown(rl.KeyD) && rl.IsKeyPressed(rl.KeyLeftShift) {
		if e.Player.Stamina >= 25 {
			e.Player.Position.X += e.Player.Speed + 10
			e.Player.IsAlive = true
			e.Player.Stamina -= 25
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
		e.Music = rl.LoadMusicStream("sounds/music/GTA San Andreas Theme Song Full ! !.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()
}

var Dead = false

func (e *Engine) MonsterCollisions() {

	for i, monster := range e.Monsters {

		// porté aggro
		if monster.Position.X > e.Player.Position.X-150 &&
			monster.Position.X < e.Player.Position.X+150 &&
			monster.Position.Y > e.Player.Position.Y-150 &&
			monster.Position.Y < e.Player.Position.Y+150 {

			e.NormalTalk(monster, fmt.Sprintf("%d", monster.Health))
			if e.Player.Position.X < e.Monsters[i].Position.X+30 && e.Monsters[i].Health > 0 {
				e.Monsters[i].Position.X -= 2
			}
			if e.Player.Position.Y < e.Monsters[i].Position.Y+30 && e.Monsters[i].Health > 0 {
				e.Monsters[i].Position.Y -= 2
			}
			if e.Player.Position.Y > e.Monsters[i].Position.Y-30 && e.Monsters[i].Health > 0 {
				e.Monsters[i].Position.Y += 2
			}
			if e.Player.Position.X > e.Monsters[i].Position.X-30 && e.Monsters[i].Health > 0 {
				e.Monsters[i].Position.X += 2
			}

			// porté attaque Monstre
			if monster.Position.X > e.Player.Position.X-35 &&
				monster.Position.X < e.Player.Position.X+35 &&
				monster.Position.Y > e.Player.Position.Y-35 &&
				monster.Position.Y < e.Player.Position.Y+35 {
				//Damage to Player
				if monster.Health > 0 && !Dead {
					Dead = true
					go func() {
						for e.Monsters[i].Health > 0 && e.Player.Health > 0 {
							// Vérifier la distance entre le joueur et le monstre
							distance := math.Sqrt(math.Pow(float64(e.Player.Position.X-e.Monsters[i].Position.X), 2) + math.Pow(float64(e.Player.Position.Y-e.Monsters[i].Position.Y), 2))
							if distance > 100 {
								// Arrêter l'attaque si le joueur est trop loin
								Dead = false
								return
							}
							e.Player.Health -= monster.Damage
							time.Sleep(1 * time.Second)
						}
						Dead = false
					}()
				}
				if e.Monsters[i].Health <= 0 && e.Monsters[i].IsAlive {
					e.Monsters[i].IsAlive = false
					e.Player.Money += e.Monsters[i].Worth

				}
			}
		}
		//Porté attaque joueur
		if monster.Position.X > e.Player.Position.X-130 &&
			monster.Position.X < e.Player.Position.X+130 &&
			monster.Position.Y > e.Player.Position.Y-130 &&
			monster.Position.Y < e.Player.Position.Y+130 {
			if rl.IsKeyPressed(rl.KeyE) && e.Monsters[i].Health > 0 {
				e.Monsters[i].Health -= 10
				// Knockback
				if e.Player.Position.X > e.Monsters[i].Position.X {
					e.Monsters[i].Position.X -= 30
				}
				if e.Player.Position.X < e.Monsters[i].Position.X {
					e.Monsters[i].Position.X += 30
				}
			}
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}
