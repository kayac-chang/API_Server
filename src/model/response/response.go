package response

import "github.com/golang/protobuf/proto"

type JSON struct {
	Code  int         `json:"-"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type ProtoBuf struct {
	Code int
	Data proto.Message
}
