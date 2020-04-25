package model

import (
	"api/model/pb"
	"time"

	"github.com/golang/protobuf/ptypes"
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

// ToProto ...
func (it Order) ToProto() (*pb.Order, error) {

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
		OrderId:     it.ID,
		GameId:      it.GameID,
		UserId:      it.UserID,
		State:       it.State.PbState(),
		Bet:         uint64(it.Bet),
		Win:         uint64(it.Win),
		CreatedAt:   createAt,
		CompletedAt: completedAt,
		UpdatedAt:   updatedAt,
	}

	return &pb, nil
}

// SubOrder ...
type SubOrder struct {
	ID string `json:"sub_order_id" db:"sub_order_id"`

	OrderID string `json:"order_id" db:"order_id"`

	State State  `json:"state" db:"state"`
	Bet   uint64 `json:"bet" db:"bet"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// State ...
type State string

const (
	Pending   State = "P"
	Completed       = "C"
	Rejected        = "R"
	Issue           = "I"
)

func (it State) PbState() pb.Order_State {

	switch it {

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

func ToState(state pb.Order_State) State {

	switch state {

	case pb.Order_Pending:
		return Pending

	case pb.Order_Completed:
		return Completed

	case pb.Order_Rejected:
		return Rejected

	case pb.Order_Issue:
		return Issue
	}

	return ""
}
