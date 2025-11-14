package main

import (
	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/infra/delivery"
	"github.com/benyamin218118/todoService/interface/controller"
	"github.com/benyamin218118/todoService/usecase"
)

type App struct {
	conf         *domain.Config
	restDelivery *delivery.RestDelivery
}

func (a *App) Init(conf *domain.Config) {
	a.conf = conf

	todoUC := usecase.NewTodoUseCase(nil, nil)
	storageUC := usecase.NewStorageUseCase(nil)
	todoCtrl := controller.NewTodoController(todoUC)
	storageCtrl := controller.NewStorageController(storageUC)
	a.restDelivery = delivery.NewRestDelivery(todoCtrl, storageCtrl, a.conf)
}

func (a *App) Run() {
	err := a.restDelivery.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
