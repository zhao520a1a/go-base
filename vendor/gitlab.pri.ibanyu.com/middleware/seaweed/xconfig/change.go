package xconfig

// ChangeType ...
type ChangeType int

// ChangeEventSource ...
type ChangeEventSource int

const (
	// ADD change type:add
	ADD ChangeType = iota
	// MODIFY change type:modify
	MODIFY
	// DELETE change type:delete
	DELETE
	// Apollo change event source:apollo
	Apollo ChangeEventSource = iota
)

// String ...
func (c ChangeType) String() string {
	switch c {
	case ADD:
		return "ADD"
	case MODIFY:
		return "MODIFY"
	case DELETE:
		return "DELETE"
	}

	return "UNKOWN"
}

// ChangeEvent ...
type ChangeEvent struct {
	Source    ChangeEventSource
	Namespace string
	Changes   map[string]*Change
}

// Change ...
type Change struct {
	OldValue   string
	NewValue   string
	ChangeType ChangeType
}
