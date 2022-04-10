package currency

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	parentCtx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ctx := context.WithValue(parentCtx, "", "")
	deadline, ok := ctx.Deadline()
	fmt.Printf("deadline %v ok %v", deadline, ok)
}
