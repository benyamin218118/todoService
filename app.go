package main

import (
	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/infra/db"
	"github.com/benyamin218118/todoService/infra/delivery"
	"github.com/benyamin218118/todoService/infra/repositories"
	"github.com/benyamin218118/todoService/interface/controller"
	"github.com/benyamin218118/todoService/usecase"
)

type App struct {
	conf         *domain.Config
	restDelivery *delivery.RestDelivery
}

func (a *App) Init(conf *domain.Config) {
	a.conf = conf

	dbConn, err := db.GetConnection(conf)
	if err != nil {
		panic(err)
	}
	redisPubSub := repositories.NewRedisPubSubRepository(conf)
	storageRepo := repositories.NewS3Storage(conf, dbConn)
	todoRepo := repositories.NewTodoMySqlRepository(dbConn)
	todoUC := usecase.NewTodoUseCase(todoRepo, storageRepo, redisPubSub)
	storageUC := usecase.NewStorageUseCase(storageRepo)
	todoCtrl := controller.NewTodoController(todoUC)
	storageCtrl := controller.NewStorageController(storageUC)
	a.restDelivery = delivery.NewRestDelivery(todoCtrl, storageCtrl, conf)
}

func (a *App) Run() {
	err := a.restDelivery.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
