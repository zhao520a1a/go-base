package util

import "fmt"

func main() {
	bs := []byte("1111ff")
	ss := BytesToBinaryBytes(bs)
	fmt.Println(string(ss), len(ss))
}

func BytesToBinaryBytes(bs []byte) []byte {
	l := len(bs)
	bl := l*8 + l + 1
	buf := make([]byte, 0, bl)
	for _, b := range bs {
		buf = appendBinaryBytes(buf, b)
	}
	return buf
}

func appendBinaryBytes(bs []byte, b byte) []byte {
	var a byte
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			bs = append(bs, '0')
		default:
			bs = append(bs, '1')
		}
		b <<= 1
	}
	return bs
}
