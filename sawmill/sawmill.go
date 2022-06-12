package sawmill

const SawBlocks = 3

var price = map[int]int{
	1: -1,
	2: 3,
	3: 1,
}

// SawTrunk produces sawn trunks
func SawTrunk(trunks []int, cutAmount int, sawnBlocks *[]int) {
	if len(trunks) == 0 {
		return
	}

	nextCutAmount := SawBlocks

	if trunks[0] > cutAmount {
		trunks[0] -= cutAmount
	} else if trunks[0] == cutAmount {
		trunks = trunks[1:]
	} else {
		nextCutAmount = cutAmount - trunks[0]
		cutAmount = trunks[0]
		trunks = trunks[1:]
	}

	*sawnBlocks = append(*sawnBlocks, cutAmount)
	SawTrunk(trunks, nextCutAmount, sawnBlocks)
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Profit calculate trunks values.
func Profit(sawnTrunks []int) int {
	total := 0

	for _, t := range sawnTrunks {
		total += price[t]
	}

	return total
}
