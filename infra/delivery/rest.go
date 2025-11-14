package delivery

import (
	"fmt"

	_ "github.com/benyamin218118/todoService/docs"
	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/interface/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files)
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

	d.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (d *RestDelivery) ListenAndServe() error {
	d.registerHandlers()
	return d.engine.Run(fmt.Sprintf("%s:%d", d.conf.ListenHost, d.conf.ListenPort))
}
