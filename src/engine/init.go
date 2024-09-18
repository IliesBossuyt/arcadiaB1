package engine

import (
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
	e.InitMap("textures/map/tilesets/map.json")

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 615, Y: 1600},
		Health:    100,
		Money:     0,
		Speed:     4,
		Stamina:   100,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2000, Y: 1600},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/slime/Slime-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2100, Y: 1650},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/slime/Slime-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2100, Y: 1550},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/slime/Slime-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2300, Y: 1150},
		Health:   30,
		Damage:   7,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2350, Y: 1150},
		Health:   30,
		Damage:   7,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2750, Y: 850},
		Health:   40,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2850, Y: 850},
		Health:   40,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3500, Y: 550},
		Health:   10,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Skeleton/Skeleton-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3500, Y: 500},
		Health:   10,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Skeleton/Skeleton-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3550, Y: 500},
		Health:   10,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Skeleton/Skeleton-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3550, Y: 550},
		Health:   10,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Skeleton/Skeleton-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4000, Y: 550},
		Health:   30,
		Damage:   7,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4000, Y: 500},
		Health:   30,
		Damage:   7,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4050, Y: 500},
		Health:   10,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Skeleton/Skeleton-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4050, Y: 550},
		Health:   10,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Skeleton/Skeleton-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4700, Y: 700},
		Health:   40,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4700, Y: 650},
		Health:   40,
		Damage:   10,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4750, Y: 720},
		Health:   50,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4750, Y: 675},
		Health:   50,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4750, Y: 625},
		Health:   50,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5400, Y: 1100},
		Health:   50,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5450, Y: 1100},
		Health:   50,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5500, Y: 1100},
		Health:   50,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5450, Y: 1150},
		Health:   20,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Skeleton/Armored Skeleton-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5500, Y: 1150},
		Health:   20,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Skeleton/Armored Skeleton-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5550, Y: 1150},
		Health:   20,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    12,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Skeleton/Armored Skeleton-Idle.png"),
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
