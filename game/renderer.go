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

func (r *Renderer) Render(mapElements []DepthAwareDrawable, lights []Lightable, hudElements []Drawable) (*ebiten.Image, error) {
	r.pre()
	r.bg()
	err := r.main(mapElements)
	if err != nil {
		return nil, err
	}

	err = r.lighting(lights)
	if err != nil {
		return nil, err
	}

	err = r.hud(hudElements)
	if err != nil {
		return nil, err
	}

	r.post()

	return r.layeredImage, nil
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
func (r *Renderer) main(elements []DepthAwareDrawable) error {
	if !r.game.invHud.IsFocused() {
		return nil
	}

	var outerErr error
	sort.SliceStable(elements, func(i, j int) bool {
		iz, err := elements[i].GetZ()
		if err != nil {
			outerErr = err
		}

		jz, err := elements[j].GetZ()
		if err != nil {
			outerErr = err
		}

		return iz < jz
	})
	if outerErr != nil {
		return outerErr
	}

	for _, e := range elements {
		err := r.game.view.DrawToMap(r.mapLayer, e)
		if err != nil {
			return nil
		}
	}
	return nil
}
func (r *Renderer) lighting(elements []Lightable) error {
	r.lightingLayer.Fill(GetAmbientLight(r.game.time))

	if !r.game.invHud.IsFocused() {
		return nil
	}

	for _, e := range elements {
		err := r.game.view.DrawToLighting(r.lightingLayer, e)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Renderer) hud(elements []Drawable) error {
	for _, e := range elements {
		err := r.game.view.DrawToHUD(r.hudLayer, e)
		if err != nil {
			return err
		}
	}
	return nil
}
