package main

import (
// 	"math/rand"
	"flag"
	"fmt"
	
// 	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/plotutil"
// 	"gonum.org/v1/plot/vg"
	
	"github.com/phil0lucas/GoForCP/SC"
	"github.com/phil0lucas/GoForCP/VS"
	"github.com/montanaflynn/stats"	
)

var infile1 = flag.String("s", "../CreateData/sc3.csv", "Name of SC input file")
var infile2 = flag.String("v", "../CreateData/vs3.csv", "Name of VS input file")
var outfile = flag.String("o", "plot01.png", "Name of output file")

// One object per Arm-Vstestcd-Visitnum
type perAVV struct {
	Arm			string
	Vstestcd	string
	Visitnum	int
	Vsstresn	float64
}

func sMerge (vs []*VS.Vsrec, subjArm map[string]string) []perAVV {
	// Create a slice of perAVV objects.
	// perAVV objects have a compound key Arm-Vstestcd-Visitnum and Vsstresn values to
	// summarize into plottable points
	var vsp []perAVV
	for _, v := range vs {
		var p perAVV
		p.Arm = subjArm[v.Usubjid]
		p.Vstestcd = v.Vstestcd
		p.Visitnum = v.Visitnum
		if v.Vsstresn != nil {
			p.Vsstresn = *v.Vsstresn
		}
		if p.Arm != "" && p.Vstestcd != "" && p.Visitnum != 0 && p.Vsstresn != 0 {
			vsp = append(vsp, p)				
		}
	}
	return vsp
}

type Key struct {
	Arm			string
	Vstestcd	string
	Visitnum	int
}

func transp (p []perAVV) map[Key][]float64 {
	// Output map
	m := make(map[Key][]float64)
	
	for _, v := range p {
		var k Key
		k.Arm = v.Arm
		k.Vstestcd = v.Vstestcd
		k.Visitnum = v.Visitnum
		m[k] = append(m[k], v.Vsstresn)
	}
	return m
}

type results struct {
	key		Key
	mean	float64
}
func calcMean (m map[Key][]float64) []results {
	var sr []results
	for k, v := range m {
		var r results
		r.key = k
		r.mean, _ = stats.Mean(v)
		sr = append(sr, r)
	}
	return sr
}

// In the final output we need a graph per value of Arm
// ie one for Active and one for Placebo
type Graph struct {
	Arm		string
}

// Each graph will have 2 lines - one each for SysBP & DiaBP
type Line struct {
	graph		Graph
	line		string		
}

type Point struct {
	line		Line
	ppoint		plotter.XYs
}

func createPoints (sr []results) []Point {
	var pp []Point
	for _, v := range sr {
		var p Point
		p.line.graph.Arm = v.key.Arm
		p.line.line = v.key.Vstestcd
		fmt.Println(p.line.graph.Arm)
		fmt.Println(p.line.line)
	}
	return pp
}


/*
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
*/

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
	vsp := sMerge(vs, subjArm)
// 	for _, v := range vsp {
// 		fmt.Println(v.Arm, v.Vstestcd, v.Visitnum, v.Vsstresn)
// 	}
	
	
// Transpose from a 'record' per Arm-Vstestcd-Visitnum-Vsstresn to one
// per Arm-Vstestcd-Visitnum, aligning the Vsstresn into slices in
// order to calculate the mean Vsstresn.
	pMap := transp(vsp)
// 	for k, v := range pMap {
// 		fmt.Printf("Arm %s, VStestCD %s, Visitnum %v, Value %v", k.Arm, k.Vstestcd, k.Visitnum, v)
// 	}
	
// Calculate the mean of the Vsstresn for each Arm-Vstestcd-Visitnum
	meanVals := calcMean(pMap)
// 	fmt.Println(meanVals)



	pp := createPoints(meanVals)
	fmt.Println(pp)
	
	
/*	

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
	
	
*/	
	
}


