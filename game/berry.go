package game

import (
	"image"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

type berryPhase uint8

func (b berryPhase) GetTexture() *ebiten.Image {
	switch b {
	case 1:
		return assets.Berries1
	case 2:
		return assets.Berries2
	case 3:
		return assets.Berries3
	case 4:
		return assets.Berries4
	case 5:
		return assets.Berries5
	case 6:
		return assets.Berries6
	case 7:
		return assets.Berries7
	case 8:
		return assets.Berries8
	default:
		panic("not a valid berry phase")
	}
}

type berryProgression struct {
	NextPhase berryPhase
	Chance    float64
}

func (b berryProgression) testChance() bool {
	return rand.Intn(1000) < int(b.Chance*1000)
}

// When given a value between 0 and 1 it maps the result
// to a growth chance based on the equation 0.000047e^(10x)
//
// This equation maps 0 -> 0 and 1 -> 1 but the change is
// very steep towards the end of the time
func mapTimeToChance(time float64) float64 {
	return 0.000047 * math.Pow(math.E, time*10)
}

func (b berryPhase) CheckForProgression(time Time) (progressions []berryProgression) {
	month := time.GetMonth()

	totalMins := int(time) / TPGM
	totalHours := totalMins / MinsPerHour
	hoursPerMonth := DaysPerMonth * HoursPerDay
	hours := totalHours % hoursPerMonth
	hoursThroughMonth := float64(hours) / float64(hoursPerMonth)
	growthChance := mapTimeToChance(hoursThroughMonth)

	switch b {
	case 1:
		if month == 1 {
			p := berryProgression{
				NextPhase: 2,
				Chance:    growthChance,
			}

			progressions = append(progressions, p)
		}
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	default:
		panic("not a valid berry phase")
	}
	return
}

type Berry struct {
	phase berryPhase
	pos   image.Point
}

func NewBerry() Berry {
	return Berry{
		phase: 1,
	}
}

func (b Berry) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, b, other)
}

func (b Berry) Origin(GameContext) image.Point {
	return b.pos
}

func (b Berry) Size(GameContext) image.Point {
	return b.phase.GetTexture().Bounds().Size()
}

func (b Berry) GetHitbox(layer GameContext) []image.Rectangle {
	width := b.Size(layer).X
	height := b.Size(layer).Y
	offsetRect := b.phase.GetTexture().Bounds().Add(b.pos).Sub(image.Pt(width/2, height))
	return []image.Rectangle{offsetRect}
}

func (b Berry) DrawAt(screen *ebiten.Image, pos image.Point) {
	width := b.Size(Render).X
	height := b.Size(Render).Y

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))
	options.GeoM.Translate(-float64(width/2), -float64(height))

	screen.DrawImage(b.phase.GetTexture(), &options)
}

func (b Berry) GetZ() int {
	return 1
}

func (b *Berry) Update(time Time) {
	progression := b.phase.CheckForProgression(time)

	for _, p := range progression {
		if p.testChance() && p.NextPhase != 0 {
			b.phase = p.NextPhase
			break
		}
	}
}
