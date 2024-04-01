package main

import (
	"fmt"
	"os"

	"github.com/vsPEach/HelmCRUD/crud-service/app"
	"github.com/vsPEach/HelmCRUD/crud-service/config"
)

func main() {
	cfg, err := config.Parse(os.Getenv("CONFIG_PATH"))
	if err != nil {
		panic(fmt.Errorf("can't parse config. %w", err))
	}
	panic(app.Run(cfg))
}
