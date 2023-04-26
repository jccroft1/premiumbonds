package premiumbonds

// int64 used because the fund total is greater than int32 max

var (
	Amounts = []int64{
		1_000_000,
		100_000,
		50_000,
		25_000,
		10_000,
		5_000,
		1_000,
		500,
		100,
		50,
		25,
	}

	MarchPrizeFund = New(330_527_200)
)

type PremiumBondPrize struct {
	PrizeFund int64
	Odds      int64
	Interest  float64
}

func New(prizeFund int64) PremiumBondPrize {
	return PremiumBondPrize{
		PrizeFund: prizeFund,
		Odds:      24_000,
		Interest:  0.033,
	}
}

func (pb PremiumBondPrize) GetPrizes() map[int64]int64 {
	out := map[int64]int64{}

	// high payout
	highPayoutShare := pb.PrizeFund / 10

	out[1_000_000] = 2
	highPayoutShare -= 2 * 1_000_000

	highPayoutAmounts := []int64{
		100_000,
		50_000,
		25_000,
		10_000,
		5_000,
	}

	split := int64(int(highPayoutShare) / len(highPayoutAmounts))
	carryOver := int64(0)

	for _, amount := range highPayoutAmounts {
		prizes := ((split + carryOver) / amount)

		// TODO: This is broken
		if split+carryOver-(amount*prizes) > amount/2 {
			prizes++
		}

		out[amount] = prizes
		carryOver = split + carryOver - (amount * prizes)
	}

	// medium payout
	number := ((pb.PrizeFund / 10) + carryOver) / 2_500

	out[1_000] = number
	out[500] = number * 3

	carryOver = ((pb.PrizeFund / 10) + carryOver) - (2_500 * number)

	// low payout
	goalNumberOfPrizes := pb.GetFundTotal() / pb.Odds

	lowPayoutShare := ((pb.PrizeFund / 10) * 8) + carryOver

	otherPrizeCount := int64(0)
	for _, p := range out {
		otherPrizeCount += p
	}

	// lowPayoutShare x = 25*a + 100*b + 50*b
	// goalNumberOfPrizes y = a + b + b + z otherPrizesCount
	// this becomes
	// b = ((x/25) - y + z)/4
	b := ((lowPayoutShare / 25) - goalNumberOfPrizes + otherPrizeCount) / 4
	out[100] = b
	out[50] = b

	// calculate 25 from remainder of the fund
	out[25] = (lowPayoutShare - (150 * b)) / 25

	return out
}

func (pb PremiumBondPrize) GetFundTotal() int64 {
	return int64(float64(pb.PrizeFund) / pb.Interest * 12)
}

// SumPrizes used for verification
func SumPrizes(prizes map[int64]int64) int64 {
	total := int64(0)

	for amount, count := range prizes {
		total += amount * count
	}

	return total
}
