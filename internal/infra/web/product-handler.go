package web

import (
	"encoding/json"
	"net/http"

	usecase "github.com/bernardinorafael/go-mensageria/internal/use-case"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductsUseCase  *usecase.ListProductsUseCase
}

func NewProductHandlers(
	createProductUseCase *usecase.CreateProductUseCase,
	listProductsUseCase *usecase.ListProductsUseCase,
) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductsUseCase:  listProductsUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var body usecase.CreateProductInputDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := p.CreateProductUseCase.Execute(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}

func (p *ProductHandlers) ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := p.ListProductsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(products)
}
