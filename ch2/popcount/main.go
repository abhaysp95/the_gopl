package popcount

// PC[i] is the population count of i
var PC [256]byte

func init() {
	for i := range PC {
		PC[i] = PC[i / 2] + byte(i & 1)
	}
}

// Popcount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	return int(PC[byte(x >> (0 * 8))] +
		PC[byte(x >> (1 * 8))] +
		PC[byte(x >> (2 * 8))] +
		PC[byte(x >> (3 * 8))] +
		PC[byte(x >> (4 * 8))] +
		PC[byte(x >> (5 * 8))] +
		PC[byte(x >> (6 * 8))] +
		PC[byte(x >> (7 * 8))])
}
