package bit

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result <<= 1 // shift left result 1 bit
		if num&1 == 1 {
			result |= 1
		}
		num >>= 1 //shift right num 1 bit
	}
	return result
}
