package controller

import (
	"net/http"
	"time"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/usecase"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	uc *usecase.TodoUseCase
}

type CreateTodoItemReq struct {
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	FileID      *string   `json:"fileId"`
}

type CreateTodoItemRes struct {
	ID string `json:"id"`
}

func NewTodoController(todoUC *usecase.TodoUseCase) *TodoController {
	return &TodoController{
		uc: todoUC,
	}
}

func (ctrl *TodoController) CreateTodo(ctx *gin.Context) {
	var input CreateTodoItemReq

	if err := ctx.BindJSON(&input); err != nil {
		ResponseBadRequest(ctx)
		return
	}

	id, err := ctrl.uc.CreateTodoItem(domain.TodoItem{
		Description: input.Description,
		DueDate:     input.DueDate,
		FileID:      input.FileID,
	})
	if HandleIfError(ctx, err) {
		return
	}

	Response(ctx, http.StatusCreated, CreateTodoItemRes{
		ID: *id,
	})
}
