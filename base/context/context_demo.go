package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "foo", "bar")
	ctx1 := context.WithValue(ctx, "foo", "bar1")
	ctx = context.WithValue(context.Background(), "foo", "bar2")
	fmt.Println(ctx.Value("foo").(string))
	fmt.Println(ctx1.Value("foo").(string))
}
