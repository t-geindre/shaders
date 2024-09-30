package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)
import _ "embed"

//go:embed stars.kage
var starsShader []byte

type StarsShader struct {
	shader              *ebiten.Shader
	start               time.Time
	scroll, scrollSpeed float32
	time                float32
	shift               float32
	size                float32
	prob                float32
	color               [4]float32
}

func NewStarsShader(ss, sh, si, pr float32, color color.Color) *StarsShader {
	shr, err := ebiten.NewShader(starsShader)
	if err != nil {
		panic(err)
	}

	r, g, b, a := color.RGBA()
	col := [4]float32{
		float32(r) / 65535,
		float32(g) / 65535,
		float32(b) / 65535,
		float32(a) / 65535,
	}

	s := &StarsShader{
		shader:      shr,
		scrollSpeed: ss,
		shift:       sh,
		start:       time.Now().Add(time.Millisecond * time.Duration(sh)),
		size:        si,
		prob:        pr,
		color:       col,
	}

	s.Update()

	return s
}

func (s *StarsShader) Update() {
	elapsed := time.Since(s.start)
	//s.start = time.Now()

	s.scroll += 1.0 / 60 * -s.scrollSpeed
	s.time = float32(elapsed.Seconds())
}

func (s *StarsShader) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]any{
			"Scroll": s.scroll,
			"Time":   s.time,
			"ShiftX": s.shift,
			"Size":   s.size,
			"Prob":   s.prob,
			"Color":  s.color,
		},
	}

	b := screen.Bounds()
	screen.DrawRectShader(b.Dx(), b.Dy(), s.shader, opts)

	return
	opts.Uniforms["ShiftX"] = float32(1000)
	opts.Uniforms["Scroll"] = s.scroll * 1.5
	screen.DrawRectShader(b.Dx(), b.Dy(), s.shader, opts)
}
