package model

import (
	"api/model/pb"
	"time"

	"github.com/golang/protobuf/ptypes"
)

// State ...
type State string

const (
	Pending   State = "P"
	Completed       = "C"
	Rejected        = "R"
	Issue           = "I"
)

// Order ...
type Order struct {
	ID string `json:"order_id" db:"order_id"`

	GameID string `json:"game_id" db:"game_id"`
	UserID string `json:"user_id" db:"user_id"`

	State State   `json:"state" db:"state"`
	Bet   float64 `json:"bet" db:"bet"`
	Win   float64 `json:"win" db:"win"`

	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (it *Order) FromProto(proto pb.Order) error {

	it.ID = proto.OrderId
	it.GameID = proto.GameId
	it.UserID = proto.UserId
	it.Bet = float64(proto.Bet)
	it.Win = float64(proto.Win)
	it.SetState(proto.State)

	if proto.CreatedAt != nil {

		time, err := ptypes.Timestamp(proto.CreatedAt)
		if err != nil {
			return err
		}

		it.CreatedAt = time
	}

	if proto.CompletedAt != nil {

		time, err := ptypes.Timestamp(proto.CompletedAt)
		if err != nil {
			return err
		}

		it.CompletedAt = time
	}

	if proto.UpdatedAt != nil {

		time, err := ptypes.Timestamp(proto.UpdatedAt)
		if err != nil {
			return err
		}

		it.UpdatedAt = time
	}

	return nil
}

// ToProto ...
func (it *Order) ToProto() (*pb.Order, error) {

	createAt, err := ptypes.TimestampProto(it.CreatedAt)
	if err != nil {
		return nil, err
	}

	completedAt, err := ptypes.TimestampProto(it.CompletedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := ptypes.TimestampProto(it.UpdatedAt)
	if err != nil {
		return nil, err
	}

	pb := pb.Order{
		OrderId: it.ID,

		GameId: it.GameID,
		UserId: it.UserID,

		State: it.PbState(),
		Bet:   uint64(it.Bet),
		Win:   uint64(it.Win),

		CreatedAt:   createAt,
		CompletedAt: completedAt,
		UpdatedAt:   updatedAt,
	}

	return &pb, nil
}

func (it *Order) PbState() pb.Order_State {

	switch it.State {

	case Pending:
		return pb.Order_Pending

	case Completed:
		return pb.Order_Completed

	case Rejected:
		return pb.Order_Rejected

	case Issue:
		return pb.Order_Issue
	}

	return -1
}

func (it *Order) SetState(state pb.Order_State) {

	switch state {

	case pb.Order_Pending:
		it.State = Pending

	case pb.Order_Completed:
		it.State = Completed

	case pb.Order_Rejected:
		it.State = Rejected

	case pb.Order_Issue:
		it.State = Issue
	}
}
