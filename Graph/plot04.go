package main

import (
// 	"math/rand"
	"flag"
	"fmt"
	
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	
	"github.com/phil0lucas/GoForCP/SC"
	"github.com/phil0lucas/GoForCP/VS"
	"github.com/montanaflynn/stats"	
)

var infile1 = flag.String("s", "../CreateData/sc3.csv", "Name of SC input file")
var infile2 = flag.String("v", "../CreateData/vs3.csv", "Name of VS input file")
var outfile = flag.String("o", "plot01.png", "Name of output file")

func UsubjidByTG(sc []*SC.Subject) map[string][]string {
	m := make(map[string][]string)
	for _, v := range sc {
		if v.Arm != nil {
			m[*v.Arm] = append(m[*v.Arm], v.Usubjid)
		}
	}
	return m
}

type Key struct {
	Arm			string
	Vstestcd	string
	Visitnum	int
}

type Point struct {
	PointKey	Key
	Vsstresn	float64
}

type Points []*Point

func sPoints (vs []*VS.Vsrec, subjArm map[string]string) Points {
	// Create a slice of Point objects.
	// Point objects have a compound key Arm-Vstestcd-Visitnum and Vsstresn values to
	// summarize into plottable points
	var vsp Points
	for _, v := range vs {
		var p Point
		p.PointKey.Arm = subjArm[v.Usubjid]
		p.PointKey.Vstestcd = v.Vstestcd
		p.PointKey.Visitnum = v.Visitnum
		if v.Vsstresn != nil {
			p.Vsstresn = *v.Vsstresn
		}
		vsp = append(vsp, &p)	
	}
	return vsp
}





func reorder (p Points) map[Key][]float64 {
	// Output map
	m := make(map[Key][]float64)
	
	for _, v := range p {
		m[v.PointKey] = append(m[v.PointKey], v.Vsstresn)
	}
	
	return m
}

type plotKeys struct {
	Arm			string
	Vstestcd	string
}

type plotValues struct {
	plotid		plotKeys
	yaxis		float64
	xaxis		int
}

func plotPoints (m map[Key][]float64) []plotValues {
	var sp []plotValues
	for k, v := range m {
		var p plotValues
		p.plotid.Arm = k.Arm
		p.plotid.Vstestcd = k.Vstestcd
		p.xaxis = k.Visitnum
		p.yaxis, _ = stats.Mean(v)
		fmt.Println (p)
		sp = append(sp, p)
	}
	return sp
}

// randomPoints returns some random x, y points.
func pPoints(q []plotValues) map[plotKeys]plotter.XYs {
	pts := make(map[plotKeys]plotter.XYs)
	for i := range q {
		fmt.Println(i)
	}
	return pts
}

func main() {
	// Read the 'SC' data and dump into the slice of structs
	sc := SC.ReadSC(infile1)
	
	// Create a map of Usubjid as key and Arm as value
	// Screening failures are eliminated from the lookup map.
	subjArm := make(map[string]string)
	for _, v := range sc {
		if v.Arm != nil {
			subjArm[v.Usubjid] = *v.Arm
		}
	}
	
// 	for k, v := range subjArm {
// 		fmt.Printf("Usubjid %s, Arm %s\n", k, v)
// 	}
	
	// Read the VS data
	vs := VS.ReadVS(infile2)
	
	// Create a slice of Point objects.
	// Point objects have a compound key Arm-Vstestcd-Visitnum and Vsstresn values to
	// summarize into plottable points
	vsp := sPoints(vs, subjArm)

	// In order to compute the stats we need all the values in slices, 
	// 	for each unique value of the Key
	pMap := reorder(vsp)
	for k, v := range pMap {
		fmt.Printf("Arm %s, VStestCD %s, Visitnum %v, Value %v", k.Arm, k.Vstestcd, k.Visitnum, v)
	}
	
	// Call the stats mean over each key value
// 	for k, v := range pMap {
// 		mean, _ := stats.Mean(v)
// 		fmt.Println(k)
// 		fmt.Println(mean)
// 	}
	
	pp := plotPoints(pMap)
	fmt.Println(pp)
	
	
	q := pPoints(pp)
	fmt.Printf("%T\n", q)
	
	
	
	

	p, err := plot.New()

	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", pPoints(pp),
		"Third", pPoints(pp))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, *outfile); err != nil {
		panic(err)
	}	
	
}


