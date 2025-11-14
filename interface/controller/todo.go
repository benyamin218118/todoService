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

// CreateTodo godoc
// @Summary Create a todo item
// @Tags todo
// @Accept json
// @Produce json
// @Param request body CreateTodoItemReq true "Todo data"
// @Success 201 {object} CreateTodoItemRes
// @Router /todo [post]
func (ctrl *TodoController) CreateTodo(ctx *gin.Context) {
	var req CreateTodoItemReq
	if err := ctx.BindJSON(&req); err != nil {
		ResponseBadRequest(ctx)
		return
	}
	input := domain.TodoItem{
		Description: req.Description,
		DueDate:     req.DueDate,
		FileID:      req.FileID,
	}
	if HandleIfError(ctx, input.Validate()) {
		return
	}

	id, err := ctrl.uc.CreateTodoItem(ctx.Request.Context(), input)
	if HandleIfError(ctx, err) {
		return
	}

	Response(ctx, http.StatusCreated, CreateTodoItemRes{
		ID: *id,
	})
}
