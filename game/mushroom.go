package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

// type berryVariant int

// const (
// 	Light berryVariant = iota
// 	Medium
// 	Dark
// )

// func berryVariantFromStr(str string) (berryVariant, error) {
// 	switch str {
// 	case "light":
// 		return Light, nil
// 	case "mid":
// 		return Medium, nil
// 	case "dark":
// 		return Dark, nil
// 	default:
// 		return 0, errors.NewUnknownBerryVariantError(str)
// 	}
// }

// func (b berryVariant) String() string {
// 	switch b {
// 	case Light:
// 		return "light"
// 	case Medium:
// 		return "mid"
// 	case Dark:
// 		return "dark"
// 	default:
// 		slog.Warn(errors.NewUnknownBerryVariantError((fmt.Sprintf("%d", int(b)))).Error())
// 		return "VariantError"
// 	}
// }

// type berryPhase int

// func (b berryPhase) String() string {
// 	return fmt.Sprintf("%d", b)
// }

// const (
// 	deathChance = 0.25
// 	deathYear   = 5
// )

// func (b berryPhase) CheckForProgression(time Time, totalAge int) (progressions []berryProgression, err error) {
// 	yearsThroughLife := int(float64(totalAge) / TPGM / MinsPerHour / HoursPerDay / DaysPerMonth / MonthsPerYear)

// 	month := time.MonthsThroughYear()
// 	throughMonth := time.ThroughMonth()

// 	switch b {
// 	case 1:
// 		progressions = b.oneMonthProgression(progressions, month, 1, throughMonth)
// 	case 2:
// 		progressions = b.twoMonthProgression(progressions, month, 2, throughMonth)
// 	case 3:
// 		progressions = b.twoMonthProgression(progressions, month, 4, throughMonth)
// 	case 4:
// 		progressions = b.twoMonthProgression(progressions, month, 6, throughMonth)
// 	case 5:
// 		progressions = b.oneMonthProgression(progressions, month, 8, throughMonth)
// 	case 6:
// 		progressions = b.oneMonthProgression(progressions, month, 1, throughMonth)
// 	case 7:
// 		progressions = b.oneMonthProgression(progressions, month, 2, throughMonth, 4)
// 	case 8:
// 	default:
// 		err = errors.NewInvalidBerryPhaseError(fmt.Sprintf("%v", b))
// 		return
// 	}
// 	progressions = b.deathProgression(progressions, yearsThroughLife, month)
// 	return
// }
// func (b berryPhase) oneMonthProgression(progressions []berryProgression, month int, growthMonth int, throughMonth float64, next ...berryPhase) []berryProgression {
// 	var nextPhase berryPhase
// 	if len(next) > 0 {
// 		nextPhase = next[0]
// 	} else {
// 		nextPhase = b + 1
// 	}

// 	if month == growthMonth {
// 		p := berryProgression{
// 			NextPhase: nextPhase,
// 			Chance:    mapTimeToChance(throughMonth),
// 		}

// 		progressions = append(progressions, p)
// 	}

// 	return progressions
// }
// func (b berryPhase) twoMonthProgression(progressions []berryProgression, month int, growthStart int, throughMonth float64) []berryProgression {
// 	percentThrough := throughMonth / 2

// 	progressions = b.oneMonthProgression(progressions, month, growthStart, percentThrough)
// 	progressions = b.oneMonthProgression(progressions, month, growthStart+1, percentThrough+0.5)

// 	return progressions
// }
// func (b berryPhase) deathProgression(progressions []berryProgression, years int, month int) []berryProgression {
// 	if years == deathYear && month >= 3 || years > deathYear {
// 		p := berryProgression{
// 			NextPhase: 8,
// 			Chance:    deathChance,
// 		}

// 		progressions = append(progressions, p)
// 	}

// 	return progressions
// }

// When given a value between 0 and 1 it maps the result
// to a growth chance based on the equation 0.01e^(5x)
//
// This equation maps 0 -> 0 and 1 -> 1 but the change is
// very steep towards the end of the time
// func mapTimeToChance(time float64) float64 {
// 	return 0.01 * math.Pow(math.E, time*5)
// }

// type berryProgression struct {
// 	NextPhase berryPhase
// 	Chance    float64
// }

// func (b berryProgression) testChance() bool {
// 	return rand.Intn(1000) < int(b.Chance*1000)
// }

type Mushroom struct {
	game *Game
	// state              state.State
	pos image.Point
	// randomTickCooldown int
	// plantedTime        Time
}

func NewMushroom(game *Game, position image.Point) (*Mushroom, error) {
	created := Mushroom{
		game: game,
		// plantedTime: game.time,
		pos: position,
	}
	// stateBuilder := state.StateBuilder{}

	// stateBuilder.Add(
	// 	state.NewProperty("age", berryPhase(1).String()),
	// 	state.NewProperty("variant", berryVariant(rand.Intn(3)).String()),
	// )

	// created.state = stateBuilder.Build()

	// created.SetCooldown(true)

	return &created, nil
}

func (m Mushroom) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, m, other)
}

func (m Mushroom) Origin(GameContext) (image.Point, error) {
	return m.pos, nil
}

func (m Mushroom) Size(GameContext) (image.Point, error) {
	texture, err := m.GetTexture()
	if err != nil {
		return image.Point{}, err
	}
	return texture.Bounds().Size(), nil
}

func (m Mushroom) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	size, err := m.Size(layer)
	if err != nil {
		return nil, err
	}
	texture, err := m.GetTexture()
	if err != nil {
		return nil, err
	}

	width := size.X
	height := size.Y
	offsetRect := texture.Bounds().Add(m.pos).Sub(image.Pt(width/2, height))
	return []image.Rectangle{offsetRect}, nil
}

func (m Mushroom) DrawAt(screen *ebiten.Image, pos image.Point) error {
	size, err := m.Size(Render)
	if err != nil {
		return err
	}
	texture, err := m.GetTexture()
	if err != nil {
		return err
	}

	width := size.X
	height := size.Y

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))
	options.GeoM.Translate(-float64(width/2), -float64(height))

	screen.DrawImage(texture, &options)
	return nil
}

func (m Mushroom) GetZ() (int, error) {
	playerHitbox, err := m.game.player.GetHitbox(Render)
	if err != nil {
		return 0, err
	}

	playerFeet := playerHitbox[0].Max.Y

	hitbox, err := m.GetHitbox(Render)
	if err != nil {
		return 0, err
	}

	if playerFeet >= hitbox[0].Max.Y {
		return -1, nil
	} else {
		return 1, nil
	}
}

const (
// Berries gets ticked once every `berryTickInterval`.
// berryTickInterval = TPGM * MinsPerHour * HoursPerDay / 2
)

func (m Mushroom) GetTexture() (*ebiten.Image, error) {
	// texturePath, err := assets.MushroomStates.GetTexturePath(m.state.ToTextureKey())
	// if err != nil {
	// 	return nil, err
	// }
	return assets.Mushrooms.GetTexture("object/mushrooms/light1"), nil
}

// func (b *Berry) SetCooldown(tickOnThisInterval bool) {
// 	timeLeftInInterval := berryTickInterval - (int(b.game.time) % berryTickInterval)
// 	if tickOnThisInterval {
// 		throughThis := rand.Intn(int(float64(timeLeftInInterval) * 0.95))
// 		b.randomTickCooldown = throughThis
// 	} else {
// 		throughNext := rand.Intn(int(float64(berryTickInterval) * 0.95))
// 		b.randomTickCooldown = timeLeftInInterval + throughNext
// 	}
// }

// func (b *Berry) Update() error {
// 	b.randomTickCooldown -= args.TimeRateFlag

// 	if b.randomTickCooldown <= 0 {

// 		currentPhase, err := state.GetIntFromState[berryPhase](b.state, "age")
// 		if err != nil {
// 			return err
// 		}

// 		progression, err := currentPhase.CheckForProgression(b.game.time, int(b.game.time-b.plantedTime))
// 		if err != nil {
// 			return err
// 		}

// 		for _, p := range progression {
// 			if p.testChance() && p.NextPhase != 0 {
// 				err := b.state.UpdateValue("age", fmt.Sprint(p.NextPhase))
// 				if err != nil {
// 					return err
// 				}
// 				break
// 			}
// 		}
// 		b.SetCooldown(false)
// 	}
// 	return nil
// }
