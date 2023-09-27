package hzl

import (
	"github.com/hazelcast/hazelcast-go-client/serialization"
)

type ByteArraySerializer struct{}

func (s *ByteArraySerializer) ID() int32 {
	return 1 // Unique ID for the serializer
}

func (s *ByteArraySerializer) Read(input serialization.DataInput) interface{} {
	return input.ReadByteArray()
}

func (s *ByteArraySerializer) Write(output serialization.DataOutput, obj interface{}) {
	data := obj.([]byte)
	output.WriteInt32(int32(len(data)))
	output.WriteByteArray(data)
}
