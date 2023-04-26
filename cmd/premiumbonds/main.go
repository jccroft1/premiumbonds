package main

import (
	"fmt"

	"github.com/jccroft1/premiumbonds/v2"
	"github.com/jccroft1/premiumbonds/v2/hypergeometric"
)

const year = 12

func main() {
	fmt.Println("entries,months,earnings,interest")

	entries := []int{25, 100, 1000, 5000, 25_000, 50_000}
	months := []int{1, 3, 6, year, 2 * year, 4 * year}

	for _, entry := range entries {
		for _, month := range months {
			hg := hypergeometric.Do(premiumbonds.MarchPrizeFund, entry, month)
			fmt.Printf("%v,%v,%v,%.4f\n", entry, month, hg.Median.Total, hg.Median.Interest)
		}
	}
}
