package entity

type OrderRepositoryInterface interface {
	Save(o *Order) error
	GetTotal() (int, error)
}
