package main

import (
	"bytes"
	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/ofpiyush/protobuf-bug/simple"
)

func main() {

	source := &simple.BytesTest{
		Data: []byte{},
		BytesEmpty: &wrappers.BytesValue{
			Value: []byte{},
		},
	}
	unmarshalledFromProto := simple.BytesTest{}
	unmarshalledFromJSON := simple.BytesTest{}

	pr, err := proto.Marshal(source)
	if err != nil {
		fmt.Println("Error in marshalling:", err)
		return
	}

	err = proto.Unmarshal(pr, &unmarshalledFromProto)
	if err != nil {
		fmt.Println("Error in unmarshalling:", err)
		return
	}

	// Json marshal of source
	sourceJSON, err := JSONMarshal(source)
	if err != nil {
		fmt.Println("Error in JSON marshalling source:", err)
		return
	}
	unmarshalledFromProtoJSON, err := JSONMarshal(&unmarshalledFromProto)
	if err != nil {
		fmt.Println("Error in JSON marshalling unmarshalled proto object:", err)
		return
	}

	err = jsonpb.Unmarshal(bytes.NewBuffer(sourceJSON.Bytes()), &unmarshalledFromJSON)
	if err != nil {
		fmt.Println("Error in JSON unmarshalling from json marshalled object:", err)
		return
	}

	unmarshalledFromJSONJSON, err := JSONMarshal(&unmarshalledFromJSON)
	if err != nil {
		fmt.Println("Error in JSON marshalling unmarshalled json object:", err)
		return
	}

	// This is the original go object
	fmt.Println("Source object:")
	fmt.Printf("Data: %#v\n", source.Data)
	fmt.Printf("BytesEmpty.Value: %#v\n", unmarshalledFromProto.BytesEmpty.Value)
	fmt.Println("Json:", sourceJSON)
	fmt.Println("Nil check on Data:", nil == source.Data)
	fmt.Println("Nil check on BytesEmpty.Value:", nil == source.BytesEmpty.Value)
	fmt.Println("---------------------------")

	fmt.Println("Unmarshalled from proto object:")
	fmt.Printf("Data: %#v\n", unmarshalledFromProto.Data)
	fmt.Printf("BytesEmpty.Value: %#v\n", unmarshalledFromProto.BytesEmpty.Value)
	fmt.Println("Json:", unmarshalledFromProtoJSON)
	fmt.Println("Nil check on Data:", nil == unmarshalledFromProto.Data)
	fmt.Println("Nil check on BytesEmpty.Value:", nil == unmarshalledFromProto.BytesEmpty.Value)
	fmt.Println("---------------------------")

	fmt.Println("Unmarshalled from json object:")
	fmt.Printf("Data: %#v\n", unmarshalledFromJSON.Data)
	fmt.Printf("BytesEmpty.Value: %#v\n", unmarshalledFromJSON.BytesEmpty.Value)
	fmt.Println("Json:", unmarshalledFromJSONJSON)
	fmt.Println("Nil check on Data:", nil == unmarshalledFromJSON.Data)
	fmt.Println("Nil check on BytesEmpty.Value:", nil == unmarshalledFromJSON.BytesEmpty.Value)

}

func JSONMarshal(m proto.Message) (*bytes.Buffer, error) {
	var b bytes.Buffer
	ms := jsonpb.Marshaler{
		EmitDefaults: true,
	}

	err := ms.Marshal(&b, m)

	if err != nil {

		return nil, err
	}
	return &b, nil
}
