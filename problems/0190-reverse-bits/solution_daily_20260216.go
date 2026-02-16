package main

// reverseBits reverses the bits of a 32-bit unsigned integer by iterating
// through all 32 bit positions, extracting each bit from the input and
// placing it into the mirrored position in the result.
func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}
