package suborder

import (
	"api/model"
	"api/model/pb"
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// Parse ...
func (it Handler) Parse(reqBody io.Reader) (*model.SubOrder, error) {

	reqByte, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}

	if len(reqByte) == 0 {
		return nil, model.ErrBodyEmpty
	}

	req := pb.SubOrder{}
	if err := proto.Unmarshal(reqByte, &req); err != nil {
		return nil, err
	}

	order := model.SubOrder{}
	if err := order.FromProto(req); err != nil {
		return nil, err
	}

	return &order, nil
}
