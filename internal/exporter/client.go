package exporter

import (
	"net/http"
	"time"
)

const (
	profileUrl = "https://zenmoney.ru/api/s1/profile/"
	loginUrl = "https://zenmoney.ru/login/enter/"
	export = "https://zenmoney.ru/export/download2011/"
)

func fetch(credentials *zenmoneyCredentials) ([]byte, error) {
	cl := 	&http.Client{Timeout: 60 * time.Second}


	return nil, nil
}

