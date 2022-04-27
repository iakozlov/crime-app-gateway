package main

import (
	"context"
	"fmt"
)

const (
//configPath = "./config/config.yaml"
)

func main() {
	//TODO: rename to context and use for db connection
	_ = context.Background()
	fmt.Println("some message")
}
