package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jonathan-m-borges/golang-app1/internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	_echo()
}

func vanilla() {
	http.HandleFunc("/order", OrderHandler)
	http.ListenAndServe(":8888", nil)
}

func _chi() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", OrderHandler)
	http.ListenAndServe(":8888", r)
}

func _echo() {
	e := echo.New()
	e.GET("/order", OrderHandlerEcho)
	e.Logger.Fatal(e.Start(":8888"))
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	order, _ := entity.NewOrder("1", 10, 1)
	json.NewEncoder(w).Encode(order)
}

func OrderHandlerEcho(c echo.Context) error {
	order, err := entity.NewOrder("1", 10, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, order)
}
