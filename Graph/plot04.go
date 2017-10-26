package main

import (
// 	"math/rand"
	"flag"
	"fmt"
	
// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/plotutil"
// 	"gonum.org/v1/plot/vg"
	
	"github.com/phil0lucas/GoForCP/SC"
	"github.com/phil0lucas/GoForCP/VS"
)

var infile1 = flag.String("s", "../CreateData/sc3.csv", "Name of SC input file")
var infile2 = flag.String("v", "../CreateData/vs3.csv", "Name of VS input file")
var outfile = flag.String("o", "plot01.png", "Name of output file")

func UsubjidByTG(sc []*SC.Subject) map[string][]string {
	m := make(map[string][]string)
	for k, v := range sc {
		if *v.arm != "" {
			
		// This evaluates to true if the key is in the map
			if val, ok := m[*v.arm]; ok {
				// Add Usubjid to the slice keyed on the Arm value
			} else {
				// Add both key and first element into the slice
			}
			
		}
	}
}


func main() {
	// Read the 'SC' data and dump into the slice of structs
	sc := SC.ReadSC(infile1)
	fmt.Printf("%T\n", sc)
	
	// Read the VS data
	vs := VS.ReadVS(infile2)
	fmt.Printf("%T\n", vs)
	
	// Create a map of key TG and value slice of USUBJIDs
	m := UsubjidByTG(sc)
	
	
	
	
	
	
	
	
	
	
	
	
	
/*	
	rand.Seed(int64(0))

	p, err := plot.New()

	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, outfile); err != nil {
		panic(err)
	}
*/	
	
}

/*
// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
*/
