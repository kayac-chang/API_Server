package response

import (
	"github.com/golang/protobuf/proto"
)

type JSON struct {
	Code  int         `json:"-"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type Link struct {
	Relation string `json:"rel"`
	Method   string `json:"method"`
	Href     string `json:"href"`
}

type ProtoBuf struct {
	Code int
	Data proto.Message
}

type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
