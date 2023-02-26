package value_object

var (
	WAITING   = "waiting"
	PAID      = "paid"
	CANCELLED = "cancelled"
)

type OrderStatus struct {
	value string
}

func NewOrderStatus(value string) OrderStatus {
	return OrderStatus{value: value}
}

func (s *OrderStatus) IsWaiting() bool {
	return s.value == WAITING
}

func (s *OrderStatus) IsPaid() bool {
	return s.value == PAID
}

func (s *OrderStatus) IsCancelled() bool {
	return s.value == CANCELLED
}
