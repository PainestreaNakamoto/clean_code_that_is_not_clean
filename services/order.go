package services

import (
	"idk/domain"
)

type OrderService interface {
	Create(*domain.Order)
}
