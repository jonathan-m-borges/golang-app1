package database

import (
	"database/sql"

	"github.com/jonathan-m-borges/golang-app1/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	orderRepository := &OrderRepository{
		Db: db,
	}
	return orderRepository
}

func (r *OrderRepository) Setup() error {
	r.Db.Exec("create table orders ( id varchar(255) not null, price float not null, tax float not null, final_price float not null, primary key (id))")
	return nil
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec(
		"insert into orders (id, price, tax, final_price) values (?, ?, ?, ?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	return err
}

func (r *OrderRepository) GetTotal() (int, error) {
	var tot int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&tot)
	if err != nil {
		return 0, err
	}
	return tot, nil
}
