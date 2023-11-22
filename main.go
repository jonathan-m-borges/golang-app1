package main

import (
	"database/sql"
	"fmt"

	"github.com/jonathan-m-borges/golang-app1/internal/infra/database"
	"github.com/jonathan-m-borges/golang-app1/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	err = orderRepository.Setup()
	if err != nil {
		panic(err)
	}

	useCase := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "2",
		Price: 10,
		Tax:   1,
	}

	output, err := useCase.Execute(input)
	if err != nil {
		panic(err)
	}

	tot, err := useCase.OrderRepository.GetTotal()
	if err != nil {
		panic(err)
	}

	print("sucesso ")
	fmt.Println(output.FinalPrice)
	fmt.Println(tot)
}
