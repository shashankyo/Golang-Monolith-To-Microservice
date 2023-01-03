package application

type productsService interface {
}

type paymentsService interface {
}

type OrderService struct {
}

func NewOrderService() {

}

type PlaceOrderCommand struct {
}

func (s *OrderService) PlaceOrder(cmd PlaceOrderCommand) {

}

type MarkOrderAsPaidCommand struct {
}

func (s OrderService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}

func (s OrderService) OrderByID(id orders.ID) (orders.Order, error) {

}
