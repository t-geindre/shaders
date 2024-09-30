package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

const ww, wh = 500, 500

func main() {

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(ww, wh)
	ebiten.RunGame(NewApp([]any{
		NewStarsShader(20*2, 0, 2, .991, colornames.White),
		NewStarsShader(30*2, 1000, 3, .993, colornames.White),
		NewStarsShader(40*2, 2000, 4, .995, colornames.White),

		//NewLaserShader(),
	}))
}

type App struct {
	objs []any
}

func NewApp(objs []any) *App {
	return &App{
		objs: objs,
	}
}

func (a *App) Update() error {
	for _, obj := range a.objs {
		if u, ok := obj.(Updatable); ok {
			u.Update()
		}
	}
	return nil
}

func (a *App) Draw(s *ebiten.Image) {
	for _, obj := range a.objs {
		if d, ok := obj.(Drawable); ok {
			d.Draw(s)
		}
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

type Updatable interface {
	Update()
}

type Drawable interface {
	Draw(*ebiten.Image)
}
