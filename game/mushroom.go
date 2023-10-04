package game

import (
	"fmt"
	"image"
	"log/slog"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/args"
	"github.com/moltenwolfcub/Forest-Game/assets"
	"github.com/moltenwolfcub/Forest-Game/errors"
	"github.com/moltenwolfcub/Forest-Game/game/state"
)

type mushroomVariant int

const (
	light1 mushroomVariant = iota
	light2
	light3
)

func mushroomVariantFromStr(str string) (mushroomVariant, error) {
	switch str {
	case "light1":
		return light1, nil
	case "light2":
		return light2, nil
	case "light3":
		return light3, nil
	default:
		return 0, errors.NewUnknownMushroomVariantError(str)
	}
}

func (m mushroomVariant) String() string {
	switch m {
	case light1:
		return "light1"
	case light2:
		return "light2"
	case light3:
		return "light3"
	default:
		slog.Warn(errors.NewUnknownMushroomVariantError((fmt.Sprintf("%d", int(m)))).Error())
		return "VariantError"
	}
}

type mushroomPhase int

func (m mushroomPhase) String() string {
	return fmt.Sprintf("%d", m)
}

// const (
// 	deathChance = 0.25
// 	deathYear   = 5
// )

func (m mushroomPhase) CheckForProgression(time Time, totalAge int) (progressions []mushroomProgression, err error) {
	// yearsThroughLife := int(float64(totalAge) / TPGM / MinsPerHour / HoursPerDay / DaysPerMonth / MonthsPerYear)

	month := time.MonthsThroughYear()
	throughMonth := time.ThroughMonth()

	switch m {
	case 1:
		//in month 6-7 increase
		progressions = m.betweenMonthProgression(progressions, month, 6, throughMonth)
	case 2:
		// in month 8-1 increase
		progressions = m.betweenMonthProgression(progressions, month, 8, throughMonth)
	case 3:
		// in month 1-2 increase
		progressions = m.betweenMonthProgression(progressions, month, 1, throughMonth)
	case 4:
		// in month 2-3 back to 3
		progressions = m.betweenMonthProgression(progressions, month, 2, throughMonth, 3)
	default:
		err = errors.NewInvalidMushroomPhaseError(fmt.Sprintf("%v", m))
		return
	}
	// progressions = m.deathProgression(progressions, yearsThroughLife, month)
	return
}

func (m mushroomPhase) betweenMonthProgression(progressions []mushroomProgression, currentMonth int, growthMonth int, timeThrough float64, next ...mushroomPhase) []mushroomProgression {
	if !(currentMonth == growthMonth || currentMonth == growthMonth%8+1) {
		return progressions
	}

	var nextPhase mushroomPhase
	if len(next) > 0 {
		nextPhase = next[0]
	} else {
		nextPhase = m + 1
	}

	if currentMonth != growthMonth {
		timeThrough += 1
	}

	progression := mushroomProgression{
		NextPhase: nextPhase,
		Chance:    mapMonthPairToChance(timeThrough),
	}
	progressions = append(progressions, progression)

	return progressions
}

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

// When given a value between 0 and 2 it maps the result
// to a growth chance based on the equation 0.06e^(2x)-0.1
func mapMonthPairToChance(time float64) float64 {
	return 0.01 * math.Pow(math.E, time*5)
}

type mushroomProgression struct {
	NextPhase mushroomPhase
	Chance    float64
}

func (m mushroomProgression) testChance() bool {
	return rand.Intn(1000) < int(m.Chance*1000)
}

type Mushroom struct {
	game               *Game
	state              state.State
	pos                image.Point
	randomTickCooldown int
	// plantedTime        Time
}

func NewMushroom(game *Game, position image.Point) (*Mushroom, error) {
	created := Mushroom{
		game: game,
		// plantedTime: game.time,
		pos: position,
	}
	stateBuilder := state.StateBuilder{}

	stateBuilder.Add(
		state.NewProperty("age", mushroomPhase(1).String()),
		state.NewProperty("variant", mushroomVariant(rand.Intn(3)).String()),
	)

	created.state = stateBuilder.Build()

	created.SetCooldown(true)

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
	// Mushrooms gets ticked once every `mushroomTickInterval`.
	mushroomTickInterval = TPGM * MinsPerHour * HoursPerDay / 2
)

func (m Mushroom) GetTexture() (*ebiten.Image, error) {
	texturePath, err := assets.MushroomStates.GetTexturePath(m.state.ToTextureKey())
	if err != nil {
		return nil, err
	}
	return assets.Mushrooms.GetTexture(texturePath), nil
}

func (m *Mushroom) SetCooldown(tickOnThisInterval bool) {
	timeLeftInInterval := mushroomTickInterval - (int(m.game.time) % mushroomTickInterval)
	if tickOnThisInterval {
		throughThis := rand.Intn(int(float64(timeLeftInInterval) * 0.95))
		m.randomTickCooldown = throughThis
	} else {
		throughNext := rand.Intn(int(float64(mushroomTickInterval) * 0.95))
		m.randomTickCooldown = timeLeftInInterval + throughNext
	}
}

func (m *Mushroom) Update() error {
	m.randomTickCooldown -= args.TimeRateFlag

	if m.randomTickCooldown <= 0 {

		currentPhase, err := state.GetIntFromState[mushroomPhase](m.state, "age")
		if err != nil {
			return err
		}

		progression, err := currentPhase.CheckForProgression(m.game.time, 0) //int(m.game.time-m.plantedTime))
		if err != nil {
			return err
		}

		for _, p := range progression {
			if p.testChance() && p.NextPhase != 0 {
				err := m.state.UpdateValue("age", fmt.Sprint(p.NextPhase))
				if err != nil {
					return err
				}
				break
			}
		}
		m.SetCooldown(false)
	}
	return nil
}
