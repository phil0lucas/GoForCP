// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image/color"
	"log"
	"math/rand"
	"math"
	"errors"
// 	"testing"

	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var (
	// DefaultLineStyle is the default style for drawing
	// lines.
	DefaultLineStyle = draw.LineStyle{
		Color:    color.Black,
		Width:    vg.Points(1),
		Dashes:   []vg.Length{},
		DashOffs: 0,
	}

	// DefaultGlyphStyle is the default style used
	// for gyph marks.
	DefaultGlyphStyle = draw.GlyphStyle{
		Color:  color.Black,
		Radius: vg.Points(2.5),
		Shape:  draw.RingGlyph{},
	}
)

var (
	ErrInfinity = errors.New("Infinite data point")
	ErrNaN      = errors.New("NaN data point")
	ErrNoData   = errors.New("No data points")
)

var (
	// DefaultGridLineStyle is the default style for grid lines.
	DefaultGridLineStyle = draw.LineStyle{
		Color: color.Gray{128},
		Width: vg.Points(0.25),
	}
)

// Scatter implements the Plotter interface, drawing
// a glyph for each of a set of points.
type Scatter struct {
	// XYs is a copy of the points for this scatter.
	XYs

	// GlyphStyle is the style of the glyphs drawn
	// at each point.
	draw.GlyphStyle
}

// NewScatter returns a Scatter that uses the
// default glyph style.
func NewScatter(xys XYer) (*Scatter, error) {
	data, err := CopyXYs(xys)
	if err != nil {
		return nil, err
	}
	return &Scatter{
		XYs:        data,
		GlyphStyle: DefaultGlyphStyle,
	}, err
}

// XYs implements the XYer interface.
type XYs []struct{ X, Y float64 }

// XYer wraps the Len and XY methods.
type XYer interface {
	// Len returns the number of x, y pairs.
	Len() int

	// XY returns an x, y pair.
	XY(int) (x, y float64)
}

// CopyXYs returns an XYs that is a copy of the x and y values from
// an XYer, or an error if one of the data points contains a NaN or
// Infinity.
func CopyXYs(data XYer) (XYs, error) {
	cpy := make(XYs, data.Len())
	for i := range cpy {
		cpy[i].X, cpy[i].Y = data.XY(i)
		if err := CheckFloats(cpy[i].X, cpy[i].Y); err != nil {
			return nil, err
		}
	}
	return cpy, nil
}

func (xys XYs) Len() int {
	return len(xys)
}

// CheckFloats returns an error if any of the arguments are NaN or Infinity.
func CheckFloats(fs ...float64) error {
	for _, f := range fs {
		switch {
		case math.IsNaN(f):
			return ErrNaN
		case math.IsInf(f, 0):
			return ErrInfinity
		}
	}
	return nil
}

// Grid implements the plot.Plotter interface, drawing
// a set of grid lines at the major tick marks.
type Grid struct {
	// Vertical is the style of the vertical lines.
	Vertical draw.LineStyle

	// Horizontal is the style of the horizontal lines.
	Horizontal draw.LineStyle
}

// NewGrid returns a new grid with both vertical and
// horizontal lines using the default grid line style.
func NewGrid() *Grid {
	return &Grid{
		Vertical:   DefaultGridLineStyle,
		Horizontal: DefaultGridLineStyle,
	}
}

// ExampleScatter draws some scatter points, a line,
// and a line with points.
func ExampleScatter() {
	rnd := rand.New(rand.NewSource(1))

	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) XYs {
		pts := make(XYs, n)
		for i := range pts {
			if i == 0 {
				pts[i].X = rnd.Float64()
			} else {
				pts[i].X = pts[i-1].X + rnd.Float64()
			}
			pts[i].Y = pts[i].X + 10*rnd.Float64()
		}
		return pts
	}

	n := 15
	scatterData := randomPoints(n)
	lineData := randomPoints(n)
	linePointsData := randomPoints(n)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(NewGrid())

	s, err := NewScatter(scatterData)
	if err != nil {
		log.Panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	l, err := NewLine(lineData)
	if err != nil {
		log.Panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	lpLine, lpPoints, err := NewLinePoints(linePointsData)
	if err != nil {
		log.Panic(err)
	}
	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpPoints.Shape = draw.CircleGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A: 255}

	p.Add(s, l, lpLine, lpPoints)
	p.Legend.Add("scatter", s)
	p.Legend.Add("line", l)
	p.Legend.Add("line points", lpLine, lpPoints)

	err = p.Save(200, 200, "scatter.png")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	ExampleScatter()
}
