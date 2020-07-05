package main

import (
"context"
"fmt"
"time"
)

func main() {
	//xcontext.
	ctx := context.WithValue(context.Background(), "foo", "bar")
	ctx1 := context.WithValue(ctx, "foo", "bar1")
	//ctx = context.WithValue(ctx, "foo1", "bar1")
	//fmt.Println(ctx.Value("foo").(string))
	//fmt.Println(ctx.Value("foo1").(string))
	fmt.Println(ctx.Value("foo").(string))
	fmt.Println(ctx1.Value("foo").(string))


}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process", ctx.Err())
	}
}
