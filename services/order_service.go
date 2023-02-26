package services

import (
	"fmt"
	"idk/domain"
	"idk/repositories"
)

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return orderService{repo: repo}
}

func (self orderService) Create(entity *domain.Order) {
	if entity.Status.IsWaiting() {
		fmt.Println("Wating")
	}

	self.repo.Save(entity)
}
