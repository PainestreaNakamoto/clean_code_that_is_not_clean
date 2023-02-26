package repositories

import (
	"fmt"
	"idk/domain"
)

type OrderDB struct {
	db string
}

func NewOrderDB() OrderRepository {
	return &OrderDB{db: "d"}
}

func (self *OrderDB) Save(entity *domain.Order) {
	fmt.Println(entity.Title)
}
