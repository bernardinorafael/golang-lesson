package usecase

import "github.com/bernardinorafael/go-mensageria/internal/entity"

type CreateProductInputDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDTO struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductsUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (uc *CreateProductUseCase) Execute(input CreateProductInputDTO) (*CreateProductOutputDTO, error) {
	product := entity.CreateNewProduct(input.Name, input.Price)

	if err := uc.ProductRepository.Create(product); err != nil {
		return nil, err
	}

	return &CreateProductOutputDTO{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
