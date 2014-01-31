package jenkins

func Hash(key []byte) uint32 {

	var hash uint32 = 0

	for _, b := range key {
		hash += uint32(b)
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}

	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)

	return hash
}
