package game

import (
	"fmt"
	"image"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/args"
	"github.com/moltenwolfcub/Forest-Game/assets"
	"github.com/moltenwolfcub/Forest-Game/game/state"
)

type berryVariant int

const (
	Light berryVariant = iota
	Medium
	Dark
)

func berryVariantFromStr(str string) berryVariant {
	switch str {
	case "light":
		return Light
	case "mid":
		return Medium
	case "dark":
		return Dark
	default:
		panic(fmt.Sprintf("Unkown berryVariant: %s", str))
	}
}

func (b berryVariant) String() string {
	switch b {
	case Light:
		return "light"
	case Medium:
		return "mid"
	case Dark:
		return "dark"
	default:
		panic("Unknown berryVariant")
	}
}

type berryPhase int

func (b berryPhase) String() string {
	return fmt.Sprintf("%d", b)
}

const (
	deathChance = 0.25
	deathYear   = 5
)

func (b berryPhase) CheckForProgression(time Time, totalAge int) (progressions []berryProgression) {
	yearsThroughLife := int(float64(totalAge) / TPGM / MinsPerHour / HoursPerDay / DaysPerMonth / MonthsPerYear)

	month := time.MonthsThroughYear()
	throughMonth := time.ThroughMonth()

	switch b {
	case 1:
		progressions = b.oneMonthProgression(progressions, month, 1, throughMonth)
	case 2:
		progressions = b.twoMonthProgression(progressions, month, 2, throughMonth)
	case 3:
		progressions = b.twoMonthProgression(progressions, month, 4, throughMonth)
	case 4:
		progressions = b.twoMonthProgression(progressions, month, 6, throughMonth)
	case 5:
		progressions = b.oneMonthProgression(progressions, month, 8, throughMonth)
	case 6:
		progressions = b.oneMonthProgression(progressions, month, 1, throughMonth)
	case 7:
		progressions = b.oneMonthProgression(progressions, month, 2, throughMonth, 4)
	case 8:
	default:
		panic("not a valid berry phase")
	}
	progressions = b.deathProgression(progressions, yearsThroughLife, month)

	return
}
func (b berryPhase) oneMonthProgression(progressions []berryProgression, month int, growthMonth int, throughMonth float64, next ...berryPhase) []berryProgression {
	var nextPhase berryPhase
	if len(next) > 0 {
		nextPhase = next[0]
	} else {
		nextPhase = b + 1
	}

	if month == growthMonth {
		p := berryProgression{
			NextPhase: nextPhase,
			Chance:    mapTimeToChance(throughMonth),
		}

		progressions = append(progressions, p)
	}

	return progressions
}
func (b berryPhase) twoMonthProgression(progressions []berryProgression, month int, growthStart int, throughMonth float64) []berryProgression {
	percentThrough := throughMonth / 2

	progressions = b.oneMonthProgression(progressions, month, growthStart, percentThrough)
	progressions = b.oneMonthProgression(progressions, month, growthStart+1, percentThrough+0.5)

	return progressions
}
func (b berryPhase) deathProgression(progressions []berryProgression, years int, month int) []berryProgression {
	if years == deathYear && month >= 3 || years > deathYear {
		p := berryProgression{
			NextPhase: 8,
			Chance:    deathChance,
		}

		progressions = append(progressions, p)
	}

	return progressions
}

// When given a value between 0 and 1 it maps the result
// to a growth chance based on the equation 0.01e^(5x)
//
// This equation maps 0 -> 0 and 1 -> 1 but the change is
// very steep towards the end of the time
func mapTimeToChance(time float64) float64 {
	return 0.01 * math.Pow(math.E, time*5)
}

type berryProgression struct {
	NextPhase berryPhase
	Chance    float64
}

func (b berryProgression) testChance() bool {
	return rand.Intn(1000) < int(b.Chance*1000)
}

type Berry struct {
	state              state.State
	pos                image.Point
	randomTickCooldown int
	plantedTime        Time
}

func NewBerry(position image.Point, time Time) Berry {
	created := Berry{
		plantedTime: time,
		pos:         position,
	}
	stateBuilder := state.StateBuilder{}
	stateBuilder.Add(
		state.NewProperty("age", berryPhase(1).String()),
		state.NewProperty("variant", berryVariant(rand.Intn(3)).String()),
	)

	created.state = stateBuilder.Build()

	created.SetCooldown(time, true)

	return created
}

func (b Berry) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, b, other)
}

func (b Berry) Origin(GameContext) image.Point {
	return b.pos
}

func (b Berry) Size(GameContext) image.Point {
	return b.GetTexture().Bounds().Size()
}

func (b Berry) GetHitbox(layer GameContext) []image.Rectangle {
	width := b.Size(layer).X
	height := b.Size(layer).Y
	offsetRect := b.GetTexture().Bounds().Add(b.pos).Sub(image.Pt(width/2, height))
	return []image.Rectangle{offsetRect}
}

func (b Berry) DrawAt(screen *ebiten.Image, pos image.Point) {
	width := b.Size(Render).X
	height := b.Size(Render).Y

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))
	options.GeoM.Translate(-float64(width/2), -float64(height))

	screen.DrawImage(b.GetTexture(), &options)
}

func (b Berry) GetZ() int {
	return 1
}

const (
	// Berries gets ticked once every `berryTickInterval`.
	berryTickInterval = TPGM * MinsPerHour * HoursPerDay / 2
)

func (b Berry) GetTexture() *ebiten.Image {
	texturePath := assets.BerryStates.GetTexturePath(b.state.ToTextureKey())
	return assets.Berries.GetTexture(texturePath)
}

func (b *Berry) SetCooldown(time Time, tickOnThis bool) {
	timeLeftInInterval := berryTickInterval - (int(time) % berryTickInterval)
	if tickOnThis {
		throughThis := rand.Intn(int(float64(timeLeftInInterval) * 0.95))
		b.randomTickCooldown = throughThis
	} else {
		throughNext := rand.Intn(int(float64(berryTickInterval) * 0.95))
		b.randomTickCooldown = timeLeftInInterval + throughNext
	}
}

func (b *Berry) Update(time Time) {
	b.randomTickCooldown -= args.TimeRateFlag

	if b.randomTickCooldown <= 0 {

		currentPhase := state.GetIntFromState[berryPhase](b.state, "age")
		progression := currentPhase.CheckForProgression(time, int(time-b.plantedTime))

		for _, p := range progression {
			if p.testChance() && p.NextPhase != 0 {
				b.state.UpdateValue("age", fmt.Sprint(p.NextPhase))
				break
			}
		}
		b.SetCooldown(time, false)
	}
}
