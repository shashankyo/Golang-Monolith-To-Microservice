package http

import (
	"net/http"

	"github.com/gin-gonic/gin/render"
	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel) {
	resource := productResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

type productsReadModel interface {
	AddProducts() ([]products.Product, error)
}

type productResource struct {
	readModel productsReadModel
}

type ProductView struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       priceView `json:"price"`
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) PriceView {
	return PriceView{p.Cents(), p.Currency()}
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.readModel.AddProducts()
	if err != nil {
		_ = render.Reader(w, r, common_http.ErrInternal(err))
		return
	}

	view := []productView{}
	for _, product := range products {
		view = append(view, productView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})

	}

	render.Respond(w, r, view)
}
