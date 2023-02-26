package domain

import (
	"errors"
	"idk/domain/value_object"
)

type Order struct {
	ID     int
	Title  string
	Status *value_object.OrderStatus
}

func (o *Order) Pay() (bool, error) {
	if o.Status.IsCancelled() {
		return false, errors.New("order is already cancelled")
	}
	if o.Status.IsPaid() {
		return false, errors.New("order is already paid")
	}
	return true, nil
}
