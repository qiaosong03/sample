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
	// 将Proto转成二进制写到文件
	err := WriteProtobufToBinaryFile(add, binaryFile)
	assert.Nil(t, err)

	// 将写到文件里的二进制数据，再读出来，看跟写入的是否一致
	add2 := &pb.AddRequest{}
	err = ReadProtobufFromBinaryFile(add2, binaryFile)
	assert.Nil(t, err)
	assert.True(t, proto.Equal(add, add2))

	// 将proto转成jsons后写入文件
	jsonFile := "../tmp/mess.json"
	err = WriteProtobufToJsonFile(add, jsonFile)
	assert.Nil(t, err)
}

// message就是grpc请求时发送的数据，将数据序列化后，写到文件里
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

// 将存到文件里的二进制数据，转成Message格式
func ReadProtobufFromBinaryFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	fmt.Println(message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}
	return nil
}

// 将message转为json后，写到文件里
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

// 将proto转成json
func ProtoBufToJson(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       "",
	}
	return marshaler.MarshalToString(message)
}
