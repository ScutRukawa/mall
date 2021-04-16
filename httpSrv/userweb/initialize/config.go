package initialize

import (
	"fmt"
	"userweb/config"
	"userweb/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	configFilePrefix := "config"
	configFilename := fmt.Sprintf("./%s_pro.yaml", configFilePrefix)

	if debug := GetEnvInfo("MALL_DEBUG"); debug {
		configFilename = fmt.Sprintf("./%s_debug.yaml", configFilePrefix)
	}
	v := viper.New()
	//文件路径设置
	v.SetConfigFile(configFilename)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	global.ServerConfig = &config.ServerConfig{}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Info("配置信息:&v", global.ServerConfig)

	//动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Info("config file change", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
		zap.S().Info(global.ServerConfig)
	})
}
