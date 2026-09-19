package storage

import (
	"fmt"
	"os"
	"time"

	"zenexport/internal/config"
)

const (
	csvFileNameTempl = "exported-transaction-%s.csv"
	timeFormat       = "2006-01-02T15-04-05"
)

func Save(data []byte, cfg *config.Config) error {
	path := cfg.OutDir + "/" + fmt.Sprintf(csvFileNameTempl, time.Now().Format(timeFormat))

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create out file %s: %w", path, err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("failed to save data to %s: %w", f.Name(), err)
	}

	return nil
}
