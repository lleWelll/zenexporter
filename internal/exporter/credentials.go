package exporter

import (
	"os"
)

type zenmoneyCredentials struct {
	username string
	password string
}

func resolveCredentials(args map[string]any) zenmoneyCredentials {
	username := os.Getenv("ZENMONEY_USERNAME")
	password := os.Getenv("ZENMONEY_PASSWORD")

	if args[usernameFlag] != nil {
		val, ok := args[usernameFlag].(string)
		if ok {
			username = val
		}
	}
	if args[passwordFlag] != nil {
		val, ok := args[passwordFlag].(string)
		if ok {
			password = val
		}
	}

	if username == "" || password == "" {
		panic(ErrInvalidCredentials)
	}

	return zenmoneyCredentials{
		username: username,
		password: password,
	}
}
