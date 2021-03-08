package test

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"qiaosong03.com/simple02_grpc/pb"
	"testing"
)

func Test_01(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/mess.bin"
	add := &pb.AddRequest{
		A: 100,
		B: 200,
	}
	err := WriteProtobufToBinaryFile(add, binaryFile)
	assert.Nil(t, err)

	add2 := &pb.AddRequest{}
	err = ReadProtobufFromBinaryFile(add2, binaryFile)
	assert.Nil(t, err)
	assert.True(t, proto.Equal(add, add2))

	jsonFile := "../tmp/mess.json"
	err = WriteProtobufToJsonFile(add, jsonFile)
	assert.Nil(t, err)
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}
	return nil
}

func WriteProtobufToJsonFile(message proto.Message, filename string) error {
	data, err := ProtoBufToJson(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to json: %w", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write json to file: %w", err)
	}
	return nil
}

func ProtoBufToJson(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       "",
	}
	return marshaler.MarshalToString(message)
}
