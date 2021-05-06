package config

//线上线下配置文件隔离
type ServerConfig struct {
	ServiceName      string             `mapstructure:"name" json:"name"`
	Port             int                `mapstructure:"port" json:"port"`
	OrderSrvInfo     OrderSrvConfig     `mapstructure:"order-srv" json:"order-srv"`
	JWTInfo          JWTConfig          `mapstructure:"jwt" json:"jwt"`
	GoodsSrvInfo     GoodsSrvConfig     `mapstructure:"goods-srv" json:"goods-srv"`
	InventorySrvInfo InventorySrvConfig `mapstructure:"inventory-srv" json:"inventory-srv"`
	Redisinfo        RedisConfig        `mapstructure:"redis" json:"redis"`
	ConsulInfo       ConsulConfig       `mapstructure:"consul" json:"consul"`
	AliPayInfo       AlipayConfig       `mapstructure:"alipay" json:"alipay"`
}

type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type InventorySrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type GoodsSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
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

type AlipayConfig struct {
	AppID        string `mapstructure:"app_id" json:"app_id"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnURL    string `mapstructure:"return_url" json:"return_url"`
}
