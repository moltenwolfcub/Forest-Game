package game

import (
	"image/color"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	layeredImage  *ebiten.Image
	bgLayer       *ebiten.Image
	mapLayer      *ebiten.Image
	lightingLayer *ebiten.Image
	hudLayer      *ebiten.Image
}

func NewRenderer() Renderer {
	render := Renderer{
		layeredImage:  ebiten.NewImage(WindowWidth, WindowHeight),
		bgLayer:       ebiten.NewImage(WindowWidth, WindowHeight),
		mapLayer:      ebiten.NewImage(WindowWidth, WindowHeight),
		hudLayer:      ebiten.NewImage(WindowWidth, WindowHeight),
		lightingLayer: ebiten.NewImage(WindowWidth, WindowHeight),
	}

	return render
}

func (r *Renderer) Render(view Viewport, time Time, mapElements []DepthAwareDrawable, lights []Lightable, hudElements []Drawable) *ebiten.Image {
	r.pre()
	r.bg()
	r.main(view, mapElements)
	r.lighting(view, time, lights)
	r.hud(view, hudElements)
	r.post()

	return r.layeredImage
}

func (r *Renderer) pre() {
	r.layeredImage.Clear()
	r.bgLayer.Clear()
	r.mapLayer.Clear()
	r.lightingLayer.Clear()
	r.hudLayer.Clear()
}
func (r *Renderer) post() {
	r.layeredImage.DrawImage(r.bgLayer, nil)
	r.layeredImage.DrawImage(r.mapLayer, nil)

	options := ebiten.DrawImageOptions{}
	options.Blend.BlendOperationRGB = ebiten.BlendOperationAdd
	options.Blend.BlendFactorSourceRGB = ebiten.BlendFactorDestinationColor
	options.Blend.BlendFactorDestinationRGB = ebiten.BlendFactorZero
	r.layeredImage.DrawImage(r.lightingLayer, &options)

	r.layeredImage.DrawImage(r.hudLayer, nil)
}

func (r *Renderer) bg() {
	r.bgLayer.Fill(color.RGBA{34, 139, 34, 255})
}
func (r *Renderer) main(view Viewport, elements []DepthAwareDrawable) {
	sort.SliceStable(elements, func(i, j int) bool {
		return elements[i].GetZ() < elements[j].GetZ()
	})

	for _, e := range elements {
		view.DrawToMap(r.mapLayer, e)
	}
}
func (r *Renderer) lighting(view Viewport, time Time, elements []Lightable) {
	r.lightingLayer.Fill(r.ambientLight(48, 255, time))

	for _, e := range elements {
		view.DrawToLighting(r.lightingLayer, e)
	}
}
func (r *Renderer) hud(view Viewport, elements []Drawable) {
	for _, e := range elements {
		view.DrawToHUD(r.hudLayer, e)
	}
}

func (r Renderer) ambientLight(min float64, max float64, time Time) color.Color {
	colorPerTick := (max - min) / float64(DAYLEN/2)
	mappedLight := min + colorPerTick*float64(time.GetTimeInDay()*TPGM)
	if mappedLight > max {
		diff := mappedLight - max
		mappedLight = max - diff
	}
	intLight := uint8(mappedLight)

	return color.RGBA{intLight, intLight, intLight, 255}
}
