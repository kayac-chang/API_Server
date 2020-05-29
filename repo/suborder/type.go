package suborder

import (
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/mediocregopher/radix/v3"
)

const table = "sub_orders"

// Repo ...
type Repo struct {
	redis redis.Redis
	db    postgres.DB
}

// New ...
func New(redis redis.Redis, db postgres.DB) Repo {

	return Repo{redis, db}
}

func (it Repo) findInRedisBySubOrderID(id string) (*model.SubOrder, error) {

	res, err := it.redis.Read("HGET", table, id)
	if err != nil {
		return nil, err
	}

	order := model.SubOrder{}
	err = json.Unmarshal([]byte(res), &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (it Repo) findInDBBySubOrderID(id string) (*model.SubOrder, error) {

	if err := it.db.Ping(); err != nil {
		return nil, err
	}

	order := model.SubOrder{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE sub_order_id=$1", table)

	if err := it.db.Get(&order, sql, id); err != nil {
		return nil, err
	}

	return &order, nil
}

// FindByID ...
func (it Repo) FindByID(id string) (*model.SubOrder, error) {

	order, err := it.findInRedisBySubOrderID(id)
	if order != nil && err == nil {
		return order, nil
	}
	if err != model.ErrNotFound {
		return nil, err
	}

	order, err = it.findInDBBySubOrderID(id)
	if order != nil && err == nil {
		return order, nil
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	return nil, model.ErrNotFound
}

func (it Repo) findAllInRedisByOrderID(orderID string) ([]model.SubOrder, error) {

	ids := []string{}

	err := it.redis.Find("LRANGE", &ids, orderID, "0", "-1")
	if err != nil {
		return nil, err
	}

	subOrders := []model.SubOrder{}

	for _, subOrderID := range ids {
		subOrder, err := it.findInRedisBySubOrderID(subOrderID)

		if err != nil {
			return nil, err
		}

		subOrders = append(subOrders, *subOrder)
	}

	return subOrders, nil
}

func (it Repo) findAllInDBByOrderID(orderID string) ([]model.SubOrder, error) {

	if err := it.db.Ping(); err != nil {
		return nil, err
	}

	subOrders := []model.SubOrder{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1", table)
	if err := it.db.Get(&subOrders, sql, orderID); err != nil {
		return nil, err
	}

	return subOrders, nil
}

func (it Repo) FindAllInOrder(orderID string) ([]model.SubOrder, error) {

	orders, err := it.findAllInRedisByOrderID(orderID)
	if orders != nil && err == nil {
		return orders, nil
	}
	if err != model.ErrNotFound {
		return nil, err
	}

	orders, err = it.findAllInDBByOrderID(orderID)
	if orders != nil && err == nil {
		return orders, nil
	}
	if err != model.ErrNotFound {
		return nil, err
	}

	return nil, err
}

// Store ...
func (it Repo) Store(subOrder *model.SubOrder) error {

	data, err := json.Marshal(subOrder)
	if err != nil {
		return err
	}

	return it.redis.Write(table, func(conn radix.Conn) error {
		err := conn.Do(
			radix.Cmd(nil, "HSET", table, subOrder.ID, string(data)),
		)
		if err != nil {
			return err
		}

		pending := "pending:" + table
		err = conn.Do(
			radix.Cmd(nil, "LPUSH", pending, string(data)),
		)
		if err != nil {
			return err
		}

		err = conn.Do(
			radix.Cmd(nil, "RPUSH", subOrder.OrderID, subOrder.ID),
		)
		if err != nil {
			return err
		}

		return nil
	})
}
