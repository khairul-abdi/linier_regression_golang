package main

import (
	"image/color"
	"log"
	"strconv"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type xy struct {
	x []float64
	y []float64
}

func (d xy) Len() int {
	return len(d.x)
}

func (d xy) XY(i int) (x, y float64) {
	x = d.x[i]
	y = d.y[i]
	return
}
func main() {
	var records = [][]string{
		//       height  weight
		{"145", "43"},
		{"152", "48"},
		{"160", "50"},
		{"170", "53"},
		{"150", "43"},
		{"180", "73"},
		{"155", "47"},
		{"165", "60"},
		{"160", "57"},
		{"163", "56"},
	}

	size := len(records)
	data := xy{
		x: make([]float64, size),
		y: make([]float64, size),
	}
	for i, v := range records {
		if len(v) != 2 {
			log.Fatal("Expected two elements")
		}
		if s, err := strconv.ParseFloat(v[0], 64); err == nil {
			data.y[i] = s
		}
		if s, err := strconv.ParseFloat(v[1], 64); err == nil {
			data.x[i] = s
		}
	}

	//Linier Regression
	b, a := stat.LinearRegression(data.x, data.y, nil, false) //set data untuk linier regession
	log.Printf("%v*x+%v", a, b)
	line := plotter.NewFunction(func(x float64) float64 { return a*x + b }) //linier regession
	line.Color = color.RGBA{R: 255, A: 255}

	p := plot.New()

	p.Title.Text = "weight vs height"
	p.X.Label.Text = "weight"
	p.Y.Label.Text = "height"

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(2)

	scatter, err := plotter.NewScatter(data)
	if err != nil {
		log.Panic(err)
	}
	scatter.Color = color.RGBA{G: 255, R: 255}

	p.Add(scatter, line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./result/example_linier.png"); err != nil {
		panic(err)
	}
}
