package actor

/*
Message 是 基于 message 字段和入参 st 来生成一个新字段
Message 封装：将处理逻辑抽象为 Apply() 方法

*/

// Message needs to be implemented by actor messages
// Apply should make use of the message's fields and the current state to
//  produce a new one
type Message[St any] interface {
	Apply(St) St
}

func NewTyped[St any](initial St) Actor[Message[St]] {
	return NewFromReducer(initial, func(msg Message[St], state St) St {
		return msg.Apply(state)
	})
}
