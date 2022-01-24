package time

import (
	"context"
	"testing"
	"time"
)

func TestStartByTimeDot(t *testing.T) {
	ctx := context.Background()

	var ch chan int
	//定时任务
	go func() {
		go StartByTimeDot(ctx)
		//go defaultManage.StartByTimePeriod(ctx)
		time.Sleep(1 * time.Second)
		ch <- 1
	}()
	<-ch
}
func TestStartPeriodFun(t *testing.T) {
	ctx := context.Background()
	var ch chan int
	go StartPeriodFun(ctx, 2*time.Minute, func(fctx context.Context) (err error) {
		err = Reload(fctx)
		if err != nil {
			return
		}
		ch <- 1
		return
	})
	<-ch
}
