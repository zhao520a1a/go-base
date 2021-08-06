package enum

//枚举从 1 开始, 除非使用零值是有意义的，如:当零值是理想的默认行为时。
type SpeakerType int32

// SpeakerType_Invalid=1, SpeakerType_Staff=1, SpeakerType_Customer=2
const (
	SpeakerType_Invalid SpeakerType = iota
	SpeakerType_Staff
	SpeakerType_Customer
)

var SpeakerTypeName = map[SpeakerType]string{
	0: "invalid",
	1: "staff",
	2: "custom",
}

var SpeakerTypeValue = map[string]SpeakerType{
	"invalid": 0,
	"staff":   1,
	"custom":  2,
}
