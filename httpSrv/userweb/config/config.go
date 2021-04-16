package config

//线上线下配置文件隔离 重点
type ServerConfig struct {
	ServiceName string        `mapstructure:"name"`
	Port        int           `mapstructure:"port"`
	MysqlInfo   UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt"`
}

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}
