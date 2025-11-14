package contracts

import "io"

type IStorage interface {
	Upload(file io.Reader, filename string) (*string, error)
}
