package global

import (
	"userweb/config"
	"userweb/proto"

	ut "github.com/go-playground/universal-translator"
)

var ServerConfig *config.ServerConfig

var Trans ut.Translator

var UseSrvClient proto.UserClient
