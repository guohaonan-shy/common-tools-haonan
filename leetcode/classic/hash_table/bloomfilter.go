package hash_table

func BloomHash(s string) int64 {
	var h int64 = 5381
	for i := range s {
		h = h*33 + int64(s[i])
	}
	return h & 0xffffffff
}

func BloomFilter(s string) bool {

	h1 := BloomHash(s)
	h2 := h1 >> 16
	n := (h1 / 32) % 2

	var bitmask int64 = (1 << (h1 % 32)) | (1 << (h2 % 32))

	bf := []int64{0x02100208, 0x04200410}

	//for i := 0; i < len(bf); i++ {
	//	if (bf[i]&bitmask == bitmask) == false {
	//		return false
	//	}
	//}
	return bf[n]&bitmask == bitmask
}

func InsertInBloomFilter(s string) int64 {
	h1 := BloomHash(s)
	h2 := h1 >> 16
	n := (h1 / 32) % 2

	var bitmask int64 = (1 << (h1 % 32)) | (1 << (h2 % 32))

	bf := []int64{0x02100208, 0x04200410}

	//for i := 0; i < len(bf); i++ {
	//	if (bf[i]&bitmask == bitmask) == false {
	//		return false
	//	}
	//}
	bf[n] |= bitmask
	return bf[n]
}
