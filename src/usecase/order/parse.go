package order

import (
	"api/model"
	"api/model/pb"
	"database/sql"

	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func (it *Usecase) Parse(reqBody io.Reader) (*model.Order, error) {

	reqByte, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}

	req := pb.Order{}
	if err := proto.Unmarshal(reqByte, &req); err != nil {
		return nil, err
	}

	created, err := parseTime(req.GetCreatedAt())
	if err != nil {
		return nil, err
	}

	completed, err := parseTime(req.GetCompletedAt())
	if err != nil {
		return nil, err
	}

	order := model.Order{
		ID: req.GetOrderId(),

		GameID: req.GetGameId(),
		UserID: req.GetUserId(),

		State: model.ToState(req.GetState()),
		Bet:   req.GetBet(),

		CreatedAt:   created,
		CompletedAt: completed,
	}

	return &order, nil
}

func parseTime(ts *timestamp.Timestamp) (sql.NullTime, error) {

	if ts == nil {
		return sql.NullTime{}, nil
	}

	time, err := ptypes.Timestamp(ts)

	if err != nil {
		return sql.NullTime{}, err
	}

	return sql.NullTime{time, true}, nil
}
