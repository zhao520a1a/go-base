package slice

type List[Data any] struct {
	ch   chan Data // 用来 同步的channel
	data []Data    // 存储数据的slice
}

func (s *List[_]) Schedule() {
	// 从 channel 接收数据
	for i := range s.ch {
		s.data = append(s.data, i)
	}
}

func (s *List[_]) Close() {
	// 最后关闭 channel
	close(s.ch)
}

func (s *List[Data]) AddData(data Data) {
	// 发送数据到 channel
	s.ch <- data
}

func (s *List[Data]) BatchAddData(dataList []Data) {
	for _, data := range dataList {
		s.ch <- data
	}
}

func (s *List[Data]) GetData() []Data {
	return s.data
}

func NewList[Data any](size int, done func()) *List[Data] {
	s := &List[Data]{
		ch:   make(chan Data, size),
		data: make([]Data, 0),
	}

	go func() {
		// 并发地 append 数据到 slice
		s.Schedule()
		done()
	}()

	return s
}
