package main

import (
	"fmt"

	"github.com/jccroft1/premiumbonds/v2"
	"github.com/jccroft1/premiumbonds/v2/hypergeometric"
	"github.com/jccroft1/premiumbonds/v2/simulation"
)

func main() {
	entries := 250
	months := 12
	simulations := 100_000

	sim := simulation.Do(premiumbonds.MarchPrizeFund, months, entries, simulations)

	hg := hypergeometric.Do(premiumbonds.MarchPrizeFund, entries, months)
	fmt.Printf("Median earnings is %v with %v%% interest\n", hg.Median.Total, hg.Median.Interest)
	fmt.Printf("Mode earnings is %v with %v%% interest\n", hg.Mode.Total, hg.Mode.Interest)
	fmt.Printf("Mean earnings is %v with %.2f%% interest\n", hg.Mean.Total, hg.Mean.Interest)

	for _, amount := range premiumbonds.Amounts {
		for i := int64(0); i < 100; i++ {
			hgProb, hgExist := hg.TotalByPrize[amount][i]
			simCount, simExist := sim.TotalByPrize[amount][i]
			simProb := float64(simCount) / float64(simulations)

			fmt.Printf("%v-%v,%.4f,%v\n", amount, i, hgProb, simProb)

			if !hgExist && !simExist {
				break
			}
		}
	}
}
