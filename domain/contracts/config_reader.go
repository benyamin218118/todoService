package contracts

import "github.com/benyamin218118/todoService/domain"

type IConfigReader interface {
	Read() (*domain.Config, error)
}
