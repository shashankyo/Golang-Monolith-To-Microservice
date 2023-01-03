package http

import (
	"net/http"

	"github.com/gin-gonic/gin/render"
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"github.com/google/uuid"
)

type ordersResource struct {
	service    application.OrderService
	repository orders.Repository
}

type PostOrderRequest struct {
	ProductID orders.ProductID `json:"product_id"`
	Address   PostOrderRequest `json:"address"`
}

type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

type PostOrderResponse struct {
	OrderID string
}

type OrderPaidView struct {
	ID     string `json:"id"`
	IsPaid bool   `json:"is_paid"`
}

func AddRoutes(router *chi.Mux, service applicatiion.OrderService, repository orders.Repository) {
	resource := ordersResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

func (o ordersResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Reader(w, r, common_http.ErrBadRequest(err))
		return
	}

	cmd := application.PlaceOrderCommand{
		OrderID:   orders.ID(uuid.NewV1().string()),
		ProductID: req.ProductID,
		Address:   application.PlaceOrderCommandAddress(req.Address),
	}
	if err := o.service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, PostOrderResponse{
		OrderID: string(cmd.OrderID),
	})
}

func (o ordersResource) GetPaid(w http.Request, r *http.Request) {
	order, err := o.repository.ByID(orders.ID(chi.URLParam(r, "id")))

	if err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}
	render.Respond(w, r, OrderPaidView(string(order.ID()), order, Paid()))

}
