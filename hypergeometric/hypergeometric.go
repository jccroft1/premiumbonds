package hypergeometric

import (
	"math"

	premiumbonds "github.com/jccroft1/premiumbonds/v2"
)

type HypergeometricParams struct {
	Trials     float64 // n
	Successes  float64 // K
	Population float64 // N
}

func (p HypergeometricParams) Mean() float64 {
	return p.Trials * (p.Successes / p.Population)
}

func (p HypergeometricParams) Mode() (int64, int64) {
	top := (p.Trials + 1) * (p.Successes + 1)
	bottom := p.Population + 2
	mode := top / bottom

	return int64(math.Ceil(mode - 1)), int64(math.Floor(mode))
}

func (p HypergeometricParams) PMF(inputK float64) float64 {
	a := BinomialWithoutExp(p.Successes, inputK)
	b := BinomialWithoutExp(p.Population-p.Successes, p.Trials-inputK)
	c := BinomialWithoutExp(p.Population, p.Trials)

	return math.Exp(a + b - c)
}

// Binomial custom calculation for highly skewed distributions
// See https://en.wikipedia.org/wiki/Binomial_coefficient#In_programming_languages
func BinomialWithoutExp(n, k float64) float64 {
	a, _ := math.Lgamma(n + 1)
	b, _ := math.Lgamma(k + 1)
	c, _ := math.Lgamma(n - k + 1)
	return a - b - c
}

type Output struct {
	Mean         Average
	Mode         Average
	Median       Average
	TotalByPrize map[int64]map[int64]float64
}

type Average struct {
	Interest float64
	Total    int64
}

func Do(pb premiumbonds.PremiumBondPrize, entries int, months int) Output {
	prizes := pb.GetPrizes()

	out := Output{}
	out.TotalByPrize = map[int64]map[int64]float64{}
	for _, amount := range premiumbonds.Amounts {
		out.TotalByPrize[amount] = map[int64]float64{}
	}

	params := HypergeometricParams{
		Trials:     float64(entries * months),
		Population: float64(pb.GetFundTotal()),
	}

	meanTotal := float64(0)
	modeTotal := int64(0)
	medianTotal := int64(0)
	for _, amount := range premiumbonds.Amounts {
		params.Successes = float64(prizes[amount])

		// Mean
		meanTotal += params.Mean() * float64(amount)

		// Mode
		_, modeTop := params.Mode()
		modeTotal += int64(modeTop) * amount

		// Median
		sum := float64(0)
		median := int64(-1)
		for k := float64(0); k < params.Successes; k++ {
			prob := params.PMF(k)

			out.TotalByPrize[amount][int64(k)] = prob
			sum += prob

			if sum > 0.5 && median == -1 {
				median = int64(k)
			}
			if prob > 0.99999 || (prob < 0.0001 && k > float64(25*months)/12) {
				break
			}
		}
		medianTotal += int64(median) * amount
	}

	out.Mean = Average{
		Total:    int64(meanTotal),
		Interest: (meanTotal * 12) / float64(entries*months) * 100,
	}
	out.Mode = Average{
		Total:    modeTotal,
		Interest: float64(modeTotal*12) / float64(entries*months) * 100,
	}
	out.Median = Average{
		Total:    medianTotal,
		Interest: float64(medianTotal*12) / float64(entries*months) * 100,
	}

	return out
}
