// ships
package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type ship struct {
	Fatness int
	Number  int
}

type Ships map[string]ship

func importShips() Ships {
	type tomlConfig struct{ AllShips Ships }
	var config tomlConfig
	if _, err := toml.DecodeFile("rules.toml", &config); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return config.AllShips
}
