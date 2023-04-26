package simulation

import (
	"fmt"
	"math/rand"

	"github.com/jccroft1/premiumbonds/v2"
)

const (
	summaryText = "Done! I earned %v over %v years with %v entries. This is an average of %v per year or %v%% in interest. \n"
)

type Output struct {
	Total        int64
	TotalByPrize map[int64]map[int64]int
	years        int
	entries      int
}

func (o Output) TotalPerYear() int64 {
	return o.Total / int64(o.years)
}

func (o Output) Interest() float32 {
	return float32(o.TotalPerYear()) / float32(o.entries)
}

func (o Output) PrintSummary() {
	fmt.Printf(summaryText, o.Total, o.years, o.entries, o.TotalPerYear(), o.Interest()*100)

	for _, amount := range premiumbonds.Amounts {
		for wonPerYear, occurances := range o.TotalByPrize[amount] {
			fmt.Printf("%v, %v, %v\n", amount, wonPerYear, occurances)
		}
	}
}

func Do(pb premiumbonds.PremiumBondPrize, months int, entries int, simulations int) Output {
	out := Output{
		years:   simulations,
		entries: entries,
	}

	progressInterval := simulations / 10

	prizes := pb.GetPrizes()

	// Create cumulative version of prize totals
	prizesCum := map[int64]int64{}
	prizesCumTotal := int64(0)
	for _, amount := range premiumbonds.Amounts {
		prizesCum[amount] += prizesCumTotal + prizes[amount]
		prizesCumTotal = prizesCum[amount]
	}
	out.TotalByPrize = map[int64]map[int64]int{}
	for _, amount := range premiumbonds.Amounts {
		out.TotalByPrize[amount] = map[int64]int{}
	}

	// Simulate for given number
	for i := 0; i < simulations; i++ {
		winningsThisRun := map[int64]int64{}

		for month := 0; month < months; month++ {
			for entry := 0; entry < entries; entry++ {

				// Generate winning number!
				num := rand.Int63n(pb.GetFundTotal())

				// check this initially for speed as it's the likely case
				if num > prizesCum[25] {
					continue
				}

				for _, amount := range premiumbonds.Amounts {
					if num < prizesCum[amount] {
						out.Total += amount
						winningsThisRun[amount]++
						break
					}
				}
			}
		}

		// Now we have the number of prizes we've won, save this to the breakdown
		for _, amount := range premiumbonds.Amounts {
			count := winningsThisRun[amount]
			out.TotalByPrize[amount][count]++
		}

		if i/10 > 0 && i%progressInterval == 0 {
			fmt.Printf("%v%%\n", i*100/simulations)
		}
	}

	return out
}
