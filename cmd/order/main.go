package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jonathan-m-borges/golang-app1/internal/infra/database"
	"github.com/jonathan-m-borges/golang-app1/internal/usecase"
	"github.com/jonathan-m-borges/golang-app1/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
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

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)
	rabbitmqWorker(msgRabbitmqChannel, uc)

	go func() {
		for {
			tot, err := uc.OrderRepository.GetTotal()
			if err != nil {
				panic(err)
			}
			print("totais de orders", tot)

		}
	}()
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("starting worker rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("mensagem processada e salva no banco:", output)
	}
}
