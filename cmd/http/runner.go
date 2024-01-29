package http

import (
	"go_playground/internal/controller/http"

	"go_playground/internal/core/server"
	"go_playground/internal/core/usecase"
	customerUcase "go_playground/internal/core/usecase/customers"

	"go_playground/internal/infrastructure/config"
	"go_playground/internal/infrastructure/repository"
	"log"
)

func Runner() {
	cfg := config.Configuration()
	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}
	app, server := server.NewHttpServer(cfg)

	/* Create the repository */
	_ = repository.NewUsersRepository(db.GetConnection())
	customerRepository := repository.NewCustomersRepositoryImpl(db.GetConnection())

	/* Create the service from usecase */
	livenessPort := usecase.NewLivenessService()
	customerServices := customerUcase.NewCustomerUseCase(customerRepository)

	/* Mount primary adapter */
	routes := http.NewBaseController(server, livenessPort)
	routes.InitRouter()

	customerRoutes := http.NewCreateController(server, customerServices)
	customerRoutes.InitCustomerRoutes()

	app.Start()
}
