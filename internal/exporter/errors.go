package exporter

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid zenmoney credentials, check ZENMONEY_USERNAME or ZENMONEY_PASSWORD environment variables")
	ErrInvalidOutputDir = errors.New("invalid output dir")
)
