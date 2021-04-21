package config

//线上线下配置文件隔离
type ServerConfig struct {
	ServiceName string        `mapstructure:"name" json:"name"`
	Port        int           `mapstructure:"port" json:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	Redisinfo   RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo  RedisConfig   `mapstructure:"consul" json:"consul"`
}

type UserSrvConfig struct {
	Host    string `mapstructure:"host" json:"host"`
	Port    int    `mapstructure:"port" json:"port"`
	SrvName string `mapstructure:"srv_name" json:"srv_name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	NameSpace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataId" json:"dataid"`
	Group     string `mapstructure:"group" json:"group"`
}
