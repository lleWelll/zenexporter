package config

import "flag"

const (
	loginFlag    = "login"
	passwordFlag = "password"
	outFlag      = "out"
)

func parseFlags() map[string]string {
	login := flag.String(loginFlag, "", "login for zenmoney")
	password := flag.String(passwordFlag, "", "pasword for zenmoney")
	out := flag.String(outFlag, "", "out directory")

	flag.Parse()

	flags := map[string]string{}
	if login != nil {
		flags[loginFlag] = *login
	}
	if password != nil {
		flags[passwordFlag] = *password
	}
	if out != nil {
		flags[outFlag] = *out
	}

	return flags
}
