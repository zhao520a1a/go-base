package time

import (
	"context"
	"testing"
	"time"
)

func TestStartByTimeDot(t *testing.T) {
	ctx := context.Background()
	var ch chan int
	go StartByTimeDot(ctx, 2, func(ctx context.Context) (err error) {
		err = Reload(ctx)
		if err != nil {
			return
		}
		ch <- 1
		return
	})
	<-ch
}

func Reload(ctx context.Context) error {
	return nil
}
func TestStartPeriodFun(t *testing.T) {
	ctx := context.Background()
	var ch chan int
	go StartByTimePeriod(ctx, 2*time.Minute, func(ctx context.Context) (err error) {
		err = Reload(ctx)
		if err != nil {
			return
		}
		ch <- 1
		return
	})
	<-ch
}
