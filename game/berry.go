package game

import (
	"image"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/args"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

type berryVariant int

const (
	Light berryVariant = iota
	Medium
	Dark
)

type berryPhase uint8

func (b berryPhase) GetTexture(variant berryVariant) *ebiten.Image {
	switch b {
	case 1:
		switch variant {
		case Dark:
			return assets.BerriesDark1
		case Medium:
			return assets.BerriesMid1
		case Light:
			return assets.BerriesLight1
		default:
			panic("not a valid berry variant")
		}
	case 2:
		switch variant {
		case Dark:
			return assets.BerriesDark2
		case Medium:
			return assets.BerriesMid2
		case Light:
			return assets.BerriesLight2
		default:
			panic("not a valid berry variant")
		}
	case 3:
		switch variant {
		case Dark:
			return assets.BerriesDark3
		case Medium:
			return assets.BerriesMid3
		case Light:
			return assets.BerriesLight3
		default:
			panic("not a valid berry variant")
		}
	case 4:
		switch variant {
		case Dark:
			return assets.BerriesDark4
		case Medium:
			return assets.BerriesMid4
		case Light:
			return assets.BerriesLight4
		default:
			panic("not a valid berry variant")
		}
	case 5:
		switch variant {
		case Dark:
			return assets.BerriesDark5
		case Medium:
			return assets.BerriesMid5
		case Light:
			return assets.BerriesLight5
		default:
			panic("not a valid berry variant")
		}
	case 6:
		switch variant {
		case Dark:
			return assets.BerriesDark6
		case Medium:
			return assets.BerriesMid6
		case Light:
			return assets.BerriesLight6
		default:
			panic("not a valid berry variant")
		}
	case 7:
		switch variant {
		case Dark:
			return assets.BerriesDark7
		case Medium:
			return assets.BerriesMid7
		case Light:
			return assets.BerriesLight7
		default:
			panic("not a valid berry variant")
		}
	case 8:
		switch variant {
		case Dark:
			return assets.BerriesDark8
		case Medium:
			return assets.BerriesMid8
		case Light:
			return assets.BerriesLight8
		default:
			panic("not a valid berry variant")
		}
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
// to a growth chance based on the equation 0.01e^(5x)
//
// This equation maps 0 -> 0 and 1 -> 1 but the change is
// very steep towards the end of the time
func mapTimeToChance(time float64) float64 {
	return 0.01 * math.Pow(math.E, time*5)
}

const (
	deathChance = 0.25
	deathYear   = 5
)

func (b berryPhase) CheckForProgression(time Time, totalAge int) (progressions []berryProgression) {
	yearsThroughLife := int(float64(totalAge) / TPGM / MinsPerHour / HoursPerDay / DaysPerMonth / MonthsPerYear)

	month := time.GetMonth()

	totalMins := int(time) / TPGM
	totalHours := totalMins / MinsPerHour
	hoursPerMonth := DaysPerMonth * HoursPerDay
	hoursThroughMonth := totalHours % hoursPerMonth
	percentThroughMonth := float64(hoursThroughMonth) / float64(hoursPerMonth)

	switch b {
	case 1:
		if month == 1 {
			p := berryProgression{
				NextPhase: 2,
				Chance:    mapTimeToChance(percentThroughMonth),
			}

			progressions = append(progressions, p)
		}
	case 2:
		percentThrough := float64(hoursThroughMonth) / float64(2*hoursPerMonth)

		if month == 2 {
			p := berryProgression{
				NextPhase: 3,
				Chance:    mapTimeToChance(percentThrough),
			}

			progressions = append(progressions, p)
		} else if month == 3 {
			p := berryProgression{
				NextPhase: 3,
				Chance:    mapTimeToChance(percentThrough + 0.5),
			}

			progressions = append(progressions, p)
		}
	case 3:
		percentThrough := float64(hoursThroughMonth) / float64(2*hoursPerMonth)

		if month == 4 {
			p := berryProgression{
				NextPhase: 4,
				Chance:    mapTimeToChance(percentThrough),
			}

			progressions = append(progressions, p)
		} else if month == 5 {
			p := berryProgression{
				NextPhase: 4,
				Chance:    mapTimeToChance(percentThrough + 0.5),
			}

			progressions = append(progressions, p)
		}
	case 4:
		percentThrough := float64(hoursThroughMonth) / float64(2*hoursPerMonth)

		if month == 6 {
			p := berryProgression{
				NextPhase: 5,
				Chance:    mapTimeToChance(percentThrough),
			}

			progressions = append(progressions, p)
		} else if month == 7 {
			p := berryProgression{
				NextPhase: 5,
				Chance:    mapTimeToChance(percentThrough + 0.5),
			}

			progressions = append(progressions, p)
		}

		if yearsThroughLife == deathYear && month >= 3 || yearsThroughLife > deathYear {
			p := berryProgression{
				NextPhase: 8,
				Chance:    deathChance,
			}

			progressions = append(progressions, p)
		}

	case 5:
		if month == 8 {
			p := berryProgression{
				NextPhase: 6,
				Chance:    mapTimeToChance(percentThroughMonth),
			}

			progressions = append(progressions, p)
		}

		if yearsThroughLife == deathYear && month >= 3 || yearsThroughLife > deathYear {
			p := berryProgression{
				NextPhase: 8,
				Chance:    deathChance,
			}

			progressions = append(progressions, p)
		}
	case 6:
		if month == 1 {
			p := berryProgression{
				NextPhase: 7,
				Chance:    mapTimeToChance(percentThroughMonth),
			}

			progressions = append(progressions, p)
		}

		if yearsThroughLife == deathYear && month >= 3 || yearsThroughLife > deathYear {
			p := berryProgression{
				NextPhase: 8,
				Chance:    deathChance,
			}

			progressions = append(progressions, p)
		}
	case 7:
		if month == 2 {
			p := berryProgression{
				NextPhase: 4,
				Chance:    mapTimeToChance(percentThroughMonth),
			}

			progressions = append(progressions, p)
		}

		if yearsThroughLife == 5 && month >= deathYear || yearsThroughLife > deathYear {
			p := berryProgression{
				NextPhase: 8,
				Chance:    deathChance,
			}

			progressions = append(progressions, p)
		}
	case 8:
	default:
		panic("not a valid berry phase")
	}
	return
}

type Berry struct {
	phase              berryPhase
	variant            berryVariant
	pos                image.Point
	randomTickCooldown int
	plantedTime        Time
}

func NewBerry(time Time) Berry {
	created := Berry{
		phase:       1,
		plantedTime: time,
		variant:     berryVariant(rand.Intn(3)),
	}
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
	return b.phase.GetTexture(b.variant).Bounds().Size()
}

func (b Berry) GetHitbox(layer GameContext) []image.Rectangle {
	width := b.Size(layer).X
	height := b.Size(layer).Y
	offsetRect := b.phase.GetTexture(b.variant).Bounds().Add(b.pos).Sub(image.Pt(width/2, height))
	return []image.Rectangle{offsetRect}
}

func (b Berry) DrawAt(screen *ebiten.Image, pos image.Point) {
	width := b.Size(Render).X
	height := b.Size(Render).Y

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))
	options.GeoM.Translate(-float64(width/2), -float64(height))

	screen.DrawImage(b.phase.GetTexture(b.variant), &options)
}

func (b Berry) GetZ() int {
	return 1
}

const (
	// Berries gets ticked once every `berryTickInterval`.
	berryTickInterval = TPGM * MinsPerHour * HoursPerDay / 2
)

func (b *Berry) SetCooldown(time Time, tickOnThis bool) {
	timeLeftInInterval := berryTickInterval - (int(time) % berryTickInterval)
	if tickOnThis {
		throughThis := rand.Intn(int(float64(timeLeftInInterval) * 0.95))
		b.randomTickCooldown = throughThis
	} else {
		throughNext := rand.Intn(int(float64(berryTickInterval) * 0.95))
		b.randomTickCooldown = timeLeftInInterval + throughNext
	}
	// fmt.Println("Next", time+Time(b.randomTickCooldown))
}

func (b *Berry) Update(time Time) {
	b.randomTickCooldown -= args.TimeRateFlag

	if b.randomTickCooldown <= 0 {
		// fmt.Print("Random Tick, ")
		progression := b.phase.CheckForProgression(time, int(time-b.plantedTime))

		for _, p := range progression {
			// fmt.Println(p.Chance)
			if p.testChance() && p.NextPhase != 0 {
				b.phase = p.NextPhase
				// if p.NextPhase == 8 {
				// 	fmt.Println("Died", time)
				// }
				break
			}
		}
		b.SetCooldown(time, false)
	}
}
