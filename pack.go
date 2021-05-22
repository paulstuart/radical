package radical

func To(a, b int32) uint64 {
	a1 := uint32(a)
	b1 := uint32(b)
	return uint64(a1)<<32 | uint64(b1)
}

func signed(x uint32) int32 {
	if x|0xF0000000 != 0 {
		return int32(x)
	}
	return -int32(x)
}

func From(x uint64) (int32, int32) {
	a := uint32(x >> 32)
	b := uint32(x & 0xFFFFFFFF)
	return signed(a), signed(b)
}

func Between(from, to, is uint64) bool {
	return from <= is && is <= to
}

/*
func main() {
	list := [][2]int32{
		{1, 2},
		{-3, 4},
		{5, -6},
		{-7, -8},
	}
	for _, item := range list {
		a, b := item[0], item[1]
		fmt.Printf("check %d and %d\n", a, b)
		c := To(a, b)
		a1, b1 := From(c)
		fmt.Printf("is now %d and %d\n", a1, b1)
	}
	fmt.Println("vim-go")
}
*/
