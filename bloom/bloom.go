package bloom

type BloomFilter struct {
	Size int
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{Size: size}
}
