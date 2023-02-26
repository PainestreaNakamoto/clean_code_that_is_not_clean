package repositories

import (
	"idk/domain"
)

type OrderRepository interface {
	Save(*domain.Order)
}
