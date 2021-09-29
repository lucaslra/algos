package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) (result int) {
	for i := 0; i <= 7; i++ {
		result += int(pc[byte(x>>(i+8))])
	}
	return
}
