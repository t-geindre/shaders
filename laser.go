package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

//go:embed laser.kage
var laserShader []byte

type LaserShader struct {
	shader    *ebiten.Shader
	start     time.Time
	thickness float32
	sqps      [][2]float32
	color     [4]float32
}

func NewLaserShader() *LaserShader {
	s, err := ebiten.NewShader(laserShader)
	if err != nil {
		panic(err)
	}

	m := float32(150)

	return &LaserShader{
		shader: s,
		start:  time.Now(),
		sqps: [][2]float32{
			{ww / 2, m},
			{ww - m, wh - m},
			{m, wh - m},
		},
	}
}

func (l *LaserShader) Update() {
	elapsed := time.Since(l.start)

	f := float32(math.Pow(math.Sin(elapsed.Seconds()), 2))
	l.thickness = float32(math.Sin(elapsed.Seconds()*100)+1)*5 + 3

	l.color = [4]float32{
		float32(0xe8) / 0xff * f,
		float32(0x00) / 0xff * f,
		float32(0x55) / 0xff * f,
		.8 * f,
	}

	f = 1
	for i := 0; i < len(l.sqps); i++ {
		l.sqps[i][0] += float32(math.Sin(elapsed.Seconds())) * f
		l.sqps[i][1] += float32(math.Cos(elapsed.Seconds())) * f
		f *= -1
	}

}

func (l *LaserShader) Draw(s *ebiten.Image) {
	opts := &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]interface{}{
			"Thickness": l.thickness,
			"Color":     l.color,
		},
	}

	b := s.Bounds()

	from := l.sqps[len(l.sqps)-1]
	for i := 0; i < len(l.sqps); i++ {
		opts.Uniforms["From"] = from
		opts.Uniforms["To"] = l.sqps[i]
		from = l.sqps[i]
		s.DrawRectShader(b.Dx(), b.Dy(), l.shader, opts)
	}

}
