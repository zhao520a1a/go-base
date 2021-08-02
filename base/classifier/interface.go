package classifier

//定义接口
type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

//实现接口
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

//实现接口
func (r Rectangle) Area() float32 {
	return r.length * r.width
}
