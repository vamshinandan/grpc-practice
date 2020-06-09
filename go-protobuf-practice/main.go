package main

import (
	"fmt"

	simplepb "github.com/simplesteph/protobuf-example-go/src/simple"
)

func main() {
	fmt.Println("Hello World!!")

	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 5},
	}

	sm.Name = "I Changed Name"
	fmt.Println(sm.GetId())
}
