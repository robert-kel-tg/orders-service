package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/robertke/orders-service/pkg/config"

	"github.com/go-chi/chi"
	"github.com/robertke/orders-service/pkg/application/command"
	"github.com/robertke/orders-service/pkg/application/handler"
	"github.com/robertke/orders-service/pkg/application/response"
	"github.com/robertke/orders-service/pkg/domain"
	"github.com/robertke/orders-service/pkg/repository"

	log "github.com/sirupsen/logrus"
)

// Add env vars support
// Add CircleCI
// Push to github as orders-service

var (
	globalConfig *config.Spec
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func initConfig() {
	var (
		err error
	)
	globalConfig, err = config.LoadEnv()
	if nil != err {
		log.WithError(err).Error("Could not load configurations from environment variables")
	}
}

func main() {

	initConfig()
	initLogger()

	orderRepo := repository.NewInMemoryOrder()
	orderHandler := handler.NewCreateOrder(orderRepo)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response.RespondWith200(w, "Orders Service")
	})

	r.Get("/orders", func(w http.ResponseWriter, r *http.Request) {
		orders := orderRepo.FindAll()
		response.RespondWith200(w, response.NewOrders(orders))
	})

	r.Get("/orders/{orderID}", func(w http.ResponseWriter, r *http.Request) {
		if orderID := chi.URLParam(r, "orderID"); orderID != "" {
			order := orderRepo.FindOneByID(orderID)
			response.RespondWith200(w, response.NewOrdersFromOne(order))
		}
	})

	r.Post("/orders", func(w http.ResponseWriter, r *http.Request) {
		var (
			order domain.Order
		)
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			log.Error(err)
		}

		if err := orderHandler.Handle(
			command.NewCreateOrder(
				order.ID,
				order.Name,
				order.Description,
			),
		); err != nil {
			log.Error(err)
		}
		response.RespondWith201(w, map[string]string{"message": "successfully created"})
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", globalConfig.Port), r); err != nil {
		log.Error(err)
	}
}
