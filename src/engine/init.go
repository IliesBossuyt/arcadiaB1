package engine

import (
	"fmt"
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Mode plein ecran
	rl.ToggleBorderlessWindowed()

	// Chargement de l'icone
	icon := rl.LoadImage("textures/map/tilesets/icon.png")
	rl.SetWindowIcon(*icon)

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitDealer()
	e.InitMap("textures/map/tilesets/map.json")

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 615, Y: 1600},
		Health:    100,
		Money:     1000,
		Speed:     10,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 400, Y: 320},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})


	e.Player.Money = 0
	e.Player.Stamina = 100
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/GTA San Andreas Theme Song Full ! !.mp3")

	rl.PlayMusicStream(e.Music)
}

func (e *Engine) InitDealer() {

	e.Dealer = entity.Dealer{
		Inv:      []item.Item{},
		Name:     "yannis",
		Position: rl.NewVector2(700, 1600),
		Sprite:   rl.LoadTexture("textures/entities/dealer/Soldier-Attack03.png"),
	}

	e.Dealer.Inv = append(e.Dealer.Inv, item.Item{
		Name:         "shild",
		Price:        1,
		IsConsumable: true,
		IsEquippable: true,
	})
	e.Dealer.Inv = append(e.Dealer.Inv, item.Item{
		Name:         "Sword",
		Price:        2,
		IsConsumable: true,
		IsEquippable: true,
	})

	e.Dealer.Inv = append(e.Dealer.Inv, item.Item{
		Name:         "potion",
		Price:        3,
		IsConsumable: true,
		IsEquippable: true,
	})
	fmt.Println(e.Dealer.Inv)
}
