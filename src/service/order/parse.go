package order

import (
	"api/model"
	"api/model/pb"
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

// Parse ...
func (it Handler) Parse(reqBody io.Reader) (*model.Order, error) {

	reqByte, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}

	req := pb.Order{}
	if err := proto.Unmarshal(reqByte, &req); err != nil {
		return nil, err
	}

	order := model.Order{
		ID: req.OrderId,

		GameID: req.GameId,
		UserID: req.UserId,

		State: model.ToState(req.State),
		Bet:   float64(req.Bet),
	}

	if req.CreatedAt != nil {

		time, err := ptypes.Timestamp(req.CreatedAt)
		if err != nil {
			return nil, err
		}

		order.CreatedAt = time
	}

	if req.CompletedAt != nil {

		time, err := ptypes.Timestamp(req.CompletedAt)
		if err != nil {
			return nil, err
		}

		order.CompletedAt = time
	}

	if req.UpdatedAt != nil {

		time, err := ptypes.Timestamp(req.UpdatedAt)
		if err != nil {
			return nil, err
		}

		order.UpdatedAt = time
	}

	return &order, nil
}
