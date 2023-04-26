package premiumbonds

import (
	"reflect"
	"testing"
)

func Test_GetPrizes(t *testing.T) {
	expected := map[int64]int64{
		1_000_000: 2,
		100_000:   62,
		50_000:    124,
		25_000:    249,
		10_000:    622,
		5_000:     1242,

		1_000: 13_220,
		500:   39_660,

		100: 1_406_020,
		50:  1_406_020,
		25:  2_140_768,
	}

	prizes := MarchPrizeFund.GetPrizes()

	if !reflect.DeepEqual(prizes, expected) {
		t.Fatalf("they didn't match, wanted: \n%v\ngot: \n%v", expected, prizes)
	}
}
