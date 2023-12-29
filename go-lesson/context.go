package golesson

import (
	"context"
	"fmt"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

const (
	myKey      = contextKey("myKey")
	anotherKey = contextKey("anotherKey")
)

func doSomething(ctx context.Context) {
	fmt.Println("Do something:", ctx.Value(myKey))
	anotherCtx := context.WithValue(ctx, anotherKey, "this is another field")
	doAnotherCtx(anotherCtx)
}

func doAnotherCtx(ctx context.Context) {
	fmt.Println("Do another thing:", ctx.Value(anotherKey))
}

// func sendValueIntoContext(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, "string_text", "hello world")
// }

// func getValueFromContext(ctx context.Context) {
// 	value := ctx.Value("string_text")
// 	fmt.Printf("the value of the context is: %s\n", value)
// }

func Process() {
	// fmt.Println("start context")
	// ctx := context.Background()
	// ctx.Done()
	// ctx = sendValueIntoContext(ctx)
	// getValueFromContext(ctx)
	ctx := context.Background()
	someCtx := context.WithValue(ctx, myKey, "My name is Sam")
	doSomething(someCtx)
}
