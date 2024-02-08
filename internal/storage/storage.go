package storage

import (
	"fmt"
	"github.com/gardasar-code/go-storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (r *Storage) UploadFile(file *file.File) (*file.File, error) {
	r.files[file.Id] = file
	return file, nil
}

func (r *Storage) Upload(name string, blob []byte) (*file.File, error) {
	//return file.NewFile(name, blob)
	newFile, err := file.NewFile(name, blob)
	if err != nil {
		return nil, err
	}

	r.files[newFile.Id] = newFile
	return newFile, nil
}

func (r *Storage) GetFile(id uuid.UUID) (*file.File, error) {
	foundFile, ok := r.files[id]
	if !ok {
		return nil, fmt.Errorf("file %v not found", id)
	}
	return foundFile, nil
}
