package storage

import (
	"errors"
	"fmt"
	"os"
	"time"

	"zenexport/internal/config"
)

const (
	csvFileNameTempl = "exported-transaction-%s.csv"
)

var ErrFileDelete = errors.New("failed to delete file")

func Save(data []byte, file *os.File) error {
	_, err := file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to save data to %s: %w", file.Name(), err)
	}

	return nil
}

func Delete(file *os.File) error {
	if file == nil {
		return ErrFileDelete
	}

	_ = file.Close()
	return os.Remove(file.Name())
}

func CreateNewFile(cfg *config.Config) (*os.File, error) {
	path := cfg.OutDir + "/" + fmt.Sprintf(csvFileNameTempl, time.Now())

	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create out file %s: %w", path, err)
	}

	return f, nil
}
