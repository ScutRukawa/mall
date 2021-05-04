package config

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	NameSpace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataId" json:"dataid"`
	Group     string `mapstructure:"group" json:"group"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

//线上线下配置文件隔离
type ServerConfig struct {
	ServiceName  string             `mapstructure:"name" json:"name"`
	Redisinfo    RedisConfig        `mapstructure:"redis" json:"redis"`
	ConsulInfo   RedisConfig        `mapstructure:"consul" json:"consul"`
	MysqlInfo    MysqlConfig        `mapstructure:"mysql" json:"mysql"`
	GoodsSrv     GoodsSrvConfig     `mapstructure:"goods-srv" json:"goods-srv"`
	InventorySrv InventorySrvConfig `mapstructure:"inventory-srv" json:"inventory-srv"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type GoodsSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type InventorySrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
