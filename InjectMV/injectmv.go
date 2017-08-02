// Take a slice of values and set a percentage to missing.
// Define the MV substitute for the type.

package injectmv

import (
	//"fmt"
	"math/rand"
	"time"
)

func Injmv(inslice []int, pctmv int, mvvalue int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	//Determine the number of values to change
	npct := (len(inslice) * pctmv) / 100
	// If there are values to change, select them randomly
	// and insert the defined 'missing value' substitute.
	if npct > 0 {
		for i := 0; i < npct; i++ {
			n := rand.Intn(len(inslice))
			//fmt.Println(n)
			inslice[n] = mvvalue
		}
	}
	return inslice
}

/*
func main() {
	var tdata = []int{1, 2, 3, 8, 12, 14, 99, 345, 34, 45, 89}
	withmv := injmv(tdata, 33, -999)
	fmt.Println(withmv)
}
*/
