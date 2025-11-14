package delivery

import (
	"fmt"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/interface/controller"
	"github.com/gin-gonic/gin"
)

type RestDelivery struct {
	conf        *domain.Config
	engine      *gin.Engine
	todoCtrl    *controller.TodoController
	storageCtrl *controller.StorageController
}

func NewRestDelivery(todoCtrl *controller.TodoController, storageCtrl *controller.StorageController, conf *domain.Config) *RestDelivery {
	return &RestDelivery{
		engine:      gin.Default(),
		todoCtrl:    todoCtrl,
		storageCtrl: storageCtrl,
		conf:        conf,
	}
}

func (d *RestDelivery) registerHandlers() {
	d.engine.POST("/todo", d.todoCtrl.CreateTodo)
	d.engine.POST("/upload", d.storageCtrl.UploadFile)
}

func (d *RestDelivery) ListenAndServe() error {
	d.registerHandlers()
	return d.engine.Run(fmt.Sprintf("%s:%s", d.conf.ListenHost, d.conf.ListenPort))
}
