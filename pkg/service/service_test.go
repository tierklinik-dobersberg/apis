package service

import (
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/wellknown"
)

type cfg struct {
	BaseConfig
	MongoConfig
}

func Test() {

	instance, _ := Configure(
		wellknown.EventV1ServiceScope,
		cfg{},
	)

	_ = instance
}
