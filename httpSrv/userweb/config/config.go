package config

//线上线下配置文件隔离
type ServerConfig struct {
	ServiceName string        `mapstructure:"name"`
	Port        int           `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt"`
	Redisinfo   RedisConfig   `mapstructure:"redis"`
	ConsulInfo  RedisConfig   `mapstructure:"consul"`
}

type UserSrvConfig struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	SrvName string `mapstructure:"srv_name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
