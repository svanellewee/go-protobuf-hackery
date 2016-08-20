package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Let's get dangerous")
	fmt.Printf("SSS %s\n\n", proto.String("hello"))
}
