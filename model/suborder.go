package model

import (
	"api/model/pb"
	"time"

	"github.com/golang/protobuf/ptypes"
)

// SubOrder ...
type SubOrder struct {
	ID string `json:"sub_order_id" db:"sub_order_id"`

	OrderID string `json:"order_id" db:"order_id"`

	State State   `json:"state" db:"state"`
	Bet   float64 `json:"bet" db:"bet"`

	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
}

func (it *SubOrder) FromProto(proto pb.SubOrder) error {

	it.ID = proto.SubOrderId
	it.OrderID = proto.OrderId
	it.Bet = float64(proto.Bet)
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
func (it *SubOrder) ToProto() (*pb.SubOrder, error) {

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

	pb := pb.SubOrder{
		SubOrderId: it.ID,
		OrderId:    it.OrderID,

		State: it.PbState(),
		Bet:   uint64(it.Bet),

		CreatedAt:   createAt,
		CompletedAt: completedAt,
		UpdatedAt:   updatedAt,
	}

	return &pb, nil
}

func (it *SubOrder) PbState() pb.SubOrder_State {

	switch it.State {

	case Pending:
		return pb.SubOrder_Pending

	case Completed:
		return pb.SubOrder_Completed

	case Rejected:
		return pb.SubOrder_Rejected

	case Issue:
		return pb.SubOrder_Issue
	}

	return -1
}

func (it *SubOrder) SetState(state pb.SubOrder_State) {

	switch state {

	case pb.SubOrder_Pending:
		it.State = Pending

	case pb.SubOrder_Completed:
		it.State = Completed

	case pb.SubOrder_Rejected:
		it.State = Rejected

	case pb.SubOrder_Issue:
		it.State = Issue
	}
}
