package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/ofpiyush/protobuf-bug/simple"
)

func main() {
	source := &simple.BytesTest{Data: []byte{}}
	unmarshalled := simple.BytesTest{}

	pr, err := proto.Marshal(source)
	if err != nil {
		fmt.Println("Error in marshalling:", err)
		return
	}

	err = proto.Unmarshal(pr, &unmarshalled)
	if err != nil {
		fmt.Println("Error in unmarshalling:", err)
		return
	}
	fmt.Printf("Data in source: %#v\n", source.Data)
	fmt.Println("Nil check result in original object:", nil == source.Data)
	fmt.Printf("Data after Unmarshal: %#v\n", unmarshalled.Data)
	fmt.Println("Nil check result in unmarshalled object:", nil == unmarshalled.Data)
}
