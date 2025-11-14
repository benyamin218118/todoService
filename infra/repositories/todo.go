package repositories

import (
	"context"
	"database/sql"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
	"github.com/google/uuid"
)

type todoRepository struct {
	db *sql.DB
}

func NewTodoMySqlRepository(db *sql.DB) contracts.ITodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (t *todoRepository) Save(ctx context.Context, todo domain.TodoItem) (id *string, err error) {
	todo.ID = uuid.New().String()
	_, err = t.db.ExecContext(ctx, "INSERT INTO todo_items (id, description, due_date, file_id) VALUES (?, ?, ?, ?)", todo.ID, todo.Description, todo.DueDate, todo.FileID)
	return &todo.ID, err
}
