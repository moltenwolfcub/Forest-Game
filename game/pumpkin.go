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

type pumpkinVariant int

const (
	pumpkinLight1 pumpkinVariant = iota
)

func pumpkinVariantFromStr(str string) (pumpkinVariant, error) {
	switch str {
	case "light1":
		return pumpkinLight1, nil
	default:
		return 0, errors.NewUnknownPumpkinVariantError(str)
	}
}

func (p pumpkinVariant) String() string {
	switch p {
	case pumpkinLight1:
		return "light1"
	default:
		slog.Warn(errors.NewUnknownPumpkinVariantError((fmt.Sprintf("%d", int(p)))).Error())
		return "VariantError"
	}
}

type pumpkinPhase int

func (p pumpkinPhase) String() string {
	return fmt.Sprintf("%d", p)
}

const (
	pumpkinDeathChance = 0.12
	pumpkinDeathYear   = 4
)

func (p pumpkinPhase) CheckForProgression(time Time, totalAge int) (progressions []pumpkinProgression, err error) {
	yearsThroughLife := int(float64(totalAge) / TPGM / MinsPerHour / HoursPerDay / DaysPerMonth / MonthsPerYear)

	month := time.MonthsThroughYear()
	throughMonth := time.ThroughMonth()

	switch p {
	case 1:
		progressions = p.oneMonthProgression(progressions, month, 2, throughMonth)
	case 2:
		progressions = p.twoMonthProgression(progressions, month, 3, throughMonth)
	case 3:
		progressions = p.twoMonthProgression(progressions, month, 5, throughMonth)
	case 4:
		progressions = p.oneMonthProgression(progressions, month, 7, throughMonth, 3)
	case 5:
		return
	default:
		err = errors.NewInvalidPumpkinPhaseError(fmt.Sprintf("%v", p))
		return
	}
	progressions = p.deathProgression(progressions, yearsThroughLife, month)
	return
}
func (p pumpkinPhase) oneMonthProgression(progressions []pumpkinProgression, month int, growthMonth int, throughMonth float64, next ...pumpkinPhase) []pumpkinProgression {
	if month != growthMonth {
		return progressions
	}

	var nextPhase pumpkinPhase
	if len(next) > 0 {
		nextPhase = next[0]
	} else {
		nextPhase = p + 1
	}

	progr := pumpkinProgression{
		NextPhase: nextPhase,
		Chance:    pumpkinMapTimeToChance(throughMonth),
	}

	progressions = append(progressions, progr)

	return progressions
}
func (p pumpkinPhase) twoMonthProgression(progressions []pumpkinProgression, month int, growthStart int, throughMonth float64) []pumpkinProgression {
	if month != growthStart && month != growthStart+1 {
		return progressions
	}
	percentThrough := throughMonth / 2

	progressions = p.oneMonthProgression(progressions, month, growthStart, percentThrough)
	progressions = p.oneMonthProgression(progressions, month, growthStart+1, percentThrough+0.5)

	return progressions
}
func (p pumpkinPhase) deathProgression(progressions []pumpkinProgression, years int, month int) []pumpkinProgression {
	if years == pumpkinDeathYear && month >= 3 || years > pumpkinDeathYear {
		progr := pumpkinProgression{
			NextPhase: 5,
			Chance:    pumpkinDeathChance,
		}

		progressions = append(progressions, progr)
	}

	return progressions
}

// When given a value between 0 and 1 it maps the result
// to a growth chance based on the equation 0.01e^(5x)
//
// This equation maps 0 -> 0 and 1 -> 1 but the change is
// very steep towards the end of the time
func pumpkinMapTimeToChance(time float64) float64 {
	return 0.01 * math.Pow(math.E, time*5)
}

type pumpkinProgression struct {
	NextPhase pumpkinPhase
	Chance    float64
}

func (b pumpkinProgression) testChance() bool {
	return rand.Intn(1000) < int(b.Chance*1000)
}

type Pumpkin struct {
	game               *Game
	state              state.State
	pos                image.Point
	randomTickCooldown int
	plantedTime        Time
}

func NewPumpkin(game *Game, position image.Point) (*Pumpkin, error) {
	created := Pumpkin{
		game:        game,
		plantedTime: game.time,
		pos:         position,
	}
	stateBuilder := state.StateBuilder{}

	stateBuilder.Add(
		state.NewProperty("age", pumpkinPhase(1).String()),
		state.NewProperty("variant", pumpkinVariant(0).String()),
	)

	created.state = stateBuilder.Build()

	created.SetCooldown(true)

	return &created, nil
}

func (p Pumpkin) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, p, other)
}

func (p Pumpkin) Origin(GameContext) (image.Point, error) {
	return p.pos, nil
}

func (p Pumpkin) Size(GameContext) (image.Point, error) {
	texture, err := p.GetTexture()
	if err != nil {
		return image.Point{}, err
	}
	return texture.Bounds().Size(), nil
}

func (p Pumpkin) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	size, err := p.Size(layer)
	if err != nil {
		return nil, err
	}
	texture, err := p.GetTexture()
	if err != nil {
		return nil, err
	}

	width := size.X
	height := size.Y
	offsetRect := texture.Bounds().Add(p.pos).Sub(image.Pt(width/2, height))
	return []image.Rectangle{offsetRect}, nil
}

func (p Pumpkin) DrawAt(screen *ebiten.Image, pos image.Point) error {
	size, err := p.Size(Render)
	if err != nil {
		return err
	}
	texture, err := p.GetTexture()
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

func (p Pumpkin) GetZ() (int, error) {
	playerHitbox, err := p.game.player.GetHitbox(Render)
	if err != nil {
		return 0, err
	}

	playerFeet := playerHitbox[0].Max.Y

	hitbox, err := p.GetHitbox(Render)
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
	// Pumpkins gets ticked once every `pumpkinTickInterval`.
	pumpkinTickInterval = TPGM * MinsPerHour * HoursPerDay / 2
)

func (p Pumpkin) GetTexture() (*ebiten.Image, error) {
	texturePath, err := assets.PumpkinStates.GetTexturePath(p.state.ToTextureKey())
	if err != nil {
		return nil, err
	}
	return assets.Pumpkins.GetTexture(texturePath), nil
}

func (p *Pumpkin) SetCooldown(tickOnThisInterval bool) {
	timeLeftInInterval := pumpkinTickInterval - (int(p.game.time) % pumpkinTickInterval)
	if tickOnThisInterval {
		throughThis := rand.Intn(int(float64(timeLeftInInterval) * 0.95))
		p.randomTickCooldown = throughThis
	} else {
		throughNext := rand.Intn(int(float64(pumpkinTickInterval) * 0.95))
		p.randomTickCooldown = timeLeftInInterval + throughNext
	}
}

func (p *Pumpkin) Update() error {
	p.randomTickCooldown -= args.TimeRateFlag

	if p.randomTickCooldown <= 0 {

		currentPhase, err := state.GetIntFromState[pumpkinPhase](p.state, "age")
		if err != nil {
			return err
		}

		progression, err := currentPhase.CheckForProgression(p.game.time, int(p.game.time-p.plantedTime))
		if err != nil {
			return err
		}

		for _, progr := range progression {
			if progr.testChance() && progr.NextPhase != 0 {
				err := p.state.UpdateValue("age", fmt.Sprint(progr.NextPhase))
				if err != nil {
					return err
				}
				break
			}
		}
		p.SetCooldown(false)
	}
	return nil
}
