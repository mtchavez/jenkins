package jenkins


func one_at_a_time(key string) uint32 {

  var hash, i uint32 = 0, 0
  runes := []rune(key)

  for ; i < uint32(len(key)); i++ {
    hash += uint32(runes[i])
    hash += (hash << 10)
    hash ^= (hash >> 6)
  }

  hash += (hash << 3);
  hash ^= (hash >> 11);
  hash += (hash << 15);

  return hash
}
