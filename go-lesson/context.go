package golesson

import (
	"context"
	"fmt"
)

func sendValueIntoContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "string_text", "hello world")
}

func getValueFromContext(ctx context.Context) {
	value := ctx.Value("string_text")
	fmt.Printf("the value of the context is: %s\n", value)
}

func Process() {
	fmt.Println("start context")
	ctx := context.Background()
	ctx.Done()
	ctx = sendValueIntoContext(ctx)
	getValueFromContext(ctx)
}
