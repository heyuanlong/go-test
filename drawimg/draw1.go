package main

import "github.com/fogleman/gg"
import "math/rand"
import "fmt"

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for j := 0; j < 1; j++ {
		for i := 0; i < 360; i += 15 {
			dc.Push()
			dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
			dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
			r := rand.Float64()
			g := rand.Float64()
			b := rand.Float64()
			a := 0 + 0.5
			dc.SetRGBA(r, g, b, a)
			dc.Fill()
			dc.Pop()
		}
		dc.SavePNG(fmt.Sprintf("out_%d.png",j))
	}
	
}
