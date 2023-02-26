package main

import (
	"idk/domain"
	"idk/domain/value_object"
	"idk/repositories"
	"idk/services"
)

func main() {
	order_repo := repositories.NewOrderDB()
	order_service := services.NewOrderService(order_repo)

	order_status := value_object.NewOrderStatus("waiting")
	new_order := domain.Order{
		ID:     1,
		Title:  "Oiler",
		Status: &order_status,
	}
	order_service.Create(&new_order)
}
