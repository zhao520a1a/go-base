package rt

const (
	MinInt48 = -(1 << 47)
	MaxInt48 = +(1 << 47) - 1
)

func PackInt(v int) uint64 {
	if u := uint64(v); v < MinInt48 || v > MaxInt48 {
		panic("int48 out of range")
	} else {
		return ((u >> 63) << 47) | (u & 0x00007fffffffffff)
	}
}

func UnpackInt(v uint64) int {
	v &= 0x0000ffffffffffff
	v |= (v >> 47) * (0xffff << 48)
	return int(v)
}
