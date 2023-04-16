package golesson

import "fmt"

// Define an interface for the service.
type Greeter interface {
	Greet() string
}

// Define a struct that implements the interface.
type EnglishGreeter struct{}

func (g *EnglishGreeter) Greet() string {
	return "Hello, world!"
}

// Define a struct that depends on the service.
type GreeterClient struct {
	greeter Greeter
}

func (c *GreeterClient) Greet() string {
	return c.greeter.Greet()
}

func Main() {
	// Create an instance of the service.
	greeter := &EnglishGreeter{}

	// Create an instance of the client and inject the service.
	client := &GreeterClient{greeter}

	// Call the client method, which calls the service method.
	fmt.Println(client.Greet())
}
