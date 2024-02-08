package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	Id   uuid.UUID
	Name string
	Data []byte
}

func NewFile(name string, blob []byte) (*File, error) {
	fileId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		Id:   fileId,
		Name: name,
		Data: blob,
	}, nil
}

func (r *File) String() string {
	return fmt.Sprintf("File(%s, %v)", r.Name, r.Id)
}
