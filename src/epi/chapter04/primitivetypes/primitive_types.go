package primitivetypes

func CountBits(x uint) uint16 {
	numBits := uint16(0)
	for x != 0 {
		numBits += uint16(x & 1)
		x >>= 1
	}
	return numBits
}

// 4.1 Compute the parity of a word
func Parity(x uint64) uint16 {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return uint16(x & 0x1)
}

func SwapBits(x int64, i, j uint) int64 {
	if ((x >> i) & 1) != ((x >> j) & 1) {
		var bitMask uint64 = (1 << i) | (1 << j)
		x ^= int64(bitMask)
	}
	return x
}

// 4.7
func Power(x float64, y int) float64 {
	result := 1.0
	var power int64 = int64(y)
	if y < 0 {
		power = -power
		x = 1.0 / x
	}

	for power != 0 {
		if power&1 != 0 {
			result *= x
		}
		x *= x
		power >>= 1
	}
	return result
}
