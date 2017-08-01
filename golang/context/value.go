package main

import (
	"context"
	"fmt"
)

func main() {
	type keyType string

	f := func(ctx context.Context, k keyType) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	ctx := context.WithValue(context.Background(), keyType("key"), keyType("value"))

	f(ctx, keyType("key"))
	f(ctx, keyType("value"))
}
