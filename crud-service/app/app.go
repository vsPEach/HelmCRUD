package app

import (
	"github.com/vsPEach/HelmCRUD/crud-service/config"
	"github.com/vsPEach/HelmCRUD/crud-service/srv"
	"github.com/vsPEach/HelmCRUD/crud-service/storage"
)

func Run(config *config.Config) error {
	newStorage, err := storage.NewStorage(config.Dsn)
	if err != nil {
		return err
	}

	return srv.Run(newStorage, config)
}
