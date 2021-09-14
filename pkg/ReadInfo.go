package ScriptHelper

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type PackageConfig struct {
	Name        string
	Version     string
	Edition     string
}

type CargoConfig struct {
	Package	PackageConfig
}

func Cargo() (string, string, error) {
	var conf CargoConfig
	_, err := toml.DecodeFile("./test_file/cargo.toml", &conf)
	if err != nil {
		// handle error
		fmt.Printf("Cargo.toml %s\n", err)
	}
	return conf.Package.Name, conf.Package.Version, err
}