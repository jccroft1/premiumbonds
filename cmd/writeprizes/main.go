package main

import (
	"fmt"
	"log"
	"os"

	premiumbonds "github.com/jccroft1/premiumbonds/v2"
)

var (
	start     = premiumbonds.MarchPrizeFund.PrizeFund
	end       = start + 10_000_000
	increment = int64(83)
)

func main() {
	f, err := os.Create("out.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "Fund")
	for _, amount := range premiumbonds.Amounts {
		fmt.Fprintf(f, "|%v", amount)
	}
	fmt.Fprintln(f, "")

	for i := start; i < end; i += increment {
		prizes := premiumbonds.New(i).GetPrizes()

		fmt.Fprintf(f, "%v", i)
		for _, amount := range premiumbonds.Amounts {
			fmt.Fprintf(f, "|%v", prizes[amount])

			// log version which can be viewed in a chart
			//fmt.Fprintf(f, "|%.0f", math.Log(float64(prizes[money.New((25))])))
		}
		fmt.Fprintln(f, "")
	}
}
