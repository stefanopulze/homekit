package config

import (
	"homekit-server/internal/homekit"
	"homekit-server/internal/homekit/entity"
	"homekit-server/internal/restapi"
)

type cfg struct {
	HomeKit  homekit.ConfigOpts
	Server   restapi.ServerOpts
	Entities entity.BaseEntities
}
