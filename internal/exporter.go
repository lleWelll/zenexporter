package internal

import (
	"zenexport/internal/client"
	"zenexport/internal/config"
	"zenexport/internal/storage"
)

func Run() error {
	c := config.MustLoad()
	client := client.NewHttpClient()

	data, err := client.ExportData(c)
	if err != nil {
		return err
	}
	if err := storage.Save(data, c); err != nil {
		return err
	}

	return nil
}
