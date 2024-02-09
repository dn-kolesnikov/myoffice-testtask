package config

import (
	"flag"
	"fmt"
)

type Config struct {
	FilePath string
}

func Parse() (Config, error) {
	var config Config
	flag.StringVar(&config.FilePath, "file", "", "")
	flag.Parse()

	if config.FilePath == "" {
		return Config{}, fmt.Errorf("usage: app -file path/to/file")
	}

	return config, nil
}
