package main

import (
	"godiff/config"
	"godiff/utils"
)

func main() {

	logger := utils.GetLogger("std")
	logger.Debug("go!")

	config, err := config.LoadConfigFromFile("./config.yaml")
	if err != nil {
		logger.Error("error reading config", "err", err)
	}

	logger.Debug("config:", "cfg", config)
}
