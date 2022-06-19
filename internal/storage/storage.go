package storage

import "storage/internal/file"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return file.NewFile(filename, blob)

	// newFile, err := file.NewFile(filename,blob)

	// if err != nil {
	// 	return nil, err
	// }

	// return newFile
}
