package model

import (
	"api/model/pb"
	"database/sql"
)

type Order struct {
	ID string `json:"order_id" db:"order_id"`

	GameID string `json:"game_id" db:"game_id"`
	UserID string `json:"user_id" db:"user_id"`

	State State  `json:"state" db:"state"`
	Bet   uint64 `json:"bet" db:"bet"`
	Win   uint64 `json:"win" db:"win"`

	CreatedAt   sql.NullTime `json:"created_at" db:"created_at"`
	CompletedAt sql.NullTime `json:"completed_at" db:"completed_at,omitempty"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at"`
}

type SubOrder struct {
	ID string `json:"sub_order_id" db:"sub_order_id"`

	OrderID string `json:"order_id" db:"order_id"`

	State State  `json:"state" db:"state"`
	Bet   uint64 `json:"bet" db:"bet"`

	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}

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
