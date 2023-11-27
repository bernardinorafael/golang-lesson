package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	apachekafka "github.com/bernardinorafael/go-mensageria/internal/infra/apache-kafka"
	"github.com/bernardinorafael/go-mensageria/internal/infra/repository"
	"github.com/bernardinorafael/go-mensageria/internal/infra/web"
	usecase "github.com/bernardinorafael/go-mensageria/internal/use-case"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.CreateNewProductRepository(db)

	createProductUseCase := usecase.NewCreateProductsUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	r := chi.NewRouter()
	r.Post("products", productHandlers.CreateProductHandler)
	r.Get("products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8080", r)

	msgChannel := make(chan *kafka.Message)
	go apachekafka.Consumer([]string{"products"}, "host.docker.internal:9094", msgChannel)

	for message := range msgChannel {
		dto := usecase.CreateProductInputDTO{}
		err := json.Unmarshal(message.Value, &dto)
		if err != nil {
			continue
		}

		_, err = createProductUseCase.Execute(dto)
		if err != nil {
			panic(err)
		}
	}
}
