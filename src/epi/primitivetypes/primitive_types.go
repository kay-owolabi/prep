package primitivetypes

func CountBits(x uint) uint16 {
	numBits := uint16(0)
	for x != 0 {
		numBits += uint16(x & 1)
		x >>= 1
	}
	return numBits
}

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
