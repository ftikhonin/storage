package storage

import (
	"fmt"
	"storage/internal/file"

	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	//return file.NewFile(filename, blob)

	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}

	s.Files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileID)
	}

	return foundFile, nil
}
