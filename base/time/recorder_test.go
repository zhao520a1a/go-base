package time

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTimeCostDecorator1(t *testing.T) {

	tests := []struct {
		f func()
	}{
		{
			func() {
				fmt.Println("--- start")
				time.Sleep(3 * time.Second)
				fmt.Println("--- end")
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 方式一
			CostTimeDecorator(tt.f)()

			// 方式二
			recorder := NewTimeRecorder()
			RecordExeTime(recorder, tt.f)()
			fmt.Printf("cost time %f s", recorder.GetCost().Seconds())
		})
	}
}
