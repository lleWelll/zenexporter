package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Login    string
	Password string
	OutDir   string
}

const (
	loginEnvVar    = "ZENMONEY_LOGIN"
	passwordEnvVar = "ZENMONEY_PASSWORD"
	outDirEnvVar   = "ZENMONEY_DATA_DIR"
)

var (
	ErrInvalidCredentials = errors.New("invalid login or password")
	ErrInvalidOutDir      = errors.New("invalid out dir")
)

func MustLoad() *Config {
	flags := parseFlags()

	c := &Config{
		Login:    firstNonEmpty(flags[loginFlag], os.Getenv(loginEnvVar)),
		Password: firstNonEmpty(flags[passwordFlag], os.Getenv(passwordEnvVar)),
		OutDir:   strings.TrimSuffix(firstNonEmpty(flags[outFlag], os.Getenv(outDirEnvVar)), "/"),
	}

	c.validateConfig()
	return c
}

func (c *Config) validateConfig() {
	if c.Login == "" || c.Password == "" {
		panic(fmt.Errorf("%w: %s", ErrInvalidCredentials, fmt.Sprintf("specify --%s, --%s params or %s, %s environment variables", loginFlag, passwordFlag, loginEnvVar, passwordEnvVar)))
	}
	if c.OutDir == "" {
		panic(fmt.Errorf("%w: %s", ErrInvalidOutDir, fmt.Sprintf("specify --%s param or %s environment variable", outFlag, outDirEnvVar)))
	}
}

func firstNonEmpty(a, b string) string {
	if strings.TrimSpace(a) != "" {
		return a
	}
	if strings.TrimSpace(b) != "" {
		return b
	}

	return ""
}
