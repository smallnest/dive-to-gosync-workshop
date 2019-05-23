package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.TODO()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0001")
	ctx = context.WithValue(ctx, "key3", "0001")
	ctx = context.WithValue(ctx, "key4", "0004")

	fmt.Println(ctx.Value("key1"))
}
