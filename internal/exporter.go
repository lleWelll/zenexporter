package internal

import (
	"zenexport/internal/client"
	"zenexport/internal/config"
	"zenexport/internal/storage"
)

func Run() error {
	c := config.MustLoad()
	client := client.NewHttpClient()

	file, err := storage.CreateNewFile(c)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := client.ExportData(c)
	if err != nil {
		storage.Delete(file)
		return err
	}
	if err := storage.Save(data, file); err != nil {
		return err
	}

	return nil
}
