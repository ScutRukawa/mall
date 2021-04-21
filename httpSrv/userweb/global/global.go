package global

import (
	"userweb/config"
	"userweb/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	Trans ut.Translator

	UseSrvClient proto.UserClient

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)
