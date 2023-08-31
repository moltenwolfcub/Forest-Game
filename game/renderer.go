package game

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	game *Game

	layeredImage  *ebiten.Image
	bgLayer       *ebiten.Image
	mapLayer      *ebiten.Image
	lightingLayer *ebiten.Image
	hudLayer      *ebiten.Image
}

func NewRenderer(game *Game) Renderer {
	render := Renderer{
		game: game,

		layeredImage:  ebiten.NewImage(WindowWidth, WindowHeight),
		bgLayer:       ebiten.NewImage(WindowWidth, WindowHeight),
		mapLayer:      ebiten.NewImage(WindowWidth, WindowHeight),
		hudLayer:      ebiten.NewImage(WindowWidth, WindowHeight),
		lightingLayer: ebiten.NewImage(WindowWidth, WindowHeight),
	}

	return render
}

func (r *Renderer) Render(mapElements []DepthAwareDrawable, lights []Lightable, hudElements []Drawable) *ebiten.Image {
	r.pre()
	r.bg()
	r.main(mapElements)
	r.lighting(lights)
	r.hud(hudElements)
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
	r.bgLayer.Fill(BackgroundColor)
}
func (r *Renderer) main(elements []DepthAwareDrawable) {
	sort.SliceStable(elements, func(i, j int) bool {
		return elements[i].GetZ() < elements[j].GetZ()
	})

	for _, e := range elements {
		r.game.view.DrawToMap(r.mapLayer, e)
	}
}
func (r *Renderer) lighting(elements []Lightable) {
	r.lightingLayer.Fill(GetAmbientLight(r.game.time))

	for _, e := range elements {
		r.game.view.DrawToLighting(r.lightingLayer, e)
	}
}
func (r *Renderer) hud(elements []Drawable) {
	for _, e := range elements {
		r.game.view.DrawToHUD(r.hudLayer, e)
	}
}
