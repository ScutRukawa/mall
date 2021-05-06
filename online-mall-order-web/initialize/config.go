package initialize

import (
	"fmt"
	"online-mall-order-web/global"

	"encoding/json"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
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
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	zap.S().Info("配置信息:&v", global.NacosConfig)

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   uint64(global.NacosConfig.Port),
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.NameSpace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})
	fmt.Println(content)
	if err := json.Unmarshal([]byte(content), global.ServerConfig); err != nil {
		zap.S().Fatalf("读取nacos配置文件失败:%s", err)
	}
	fmt.Println(global.ServerConfig)

	client.ListenConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	// //动态监控变化
	// v.WatchConfig()
	// v.OnConfigChange(func(e fsnotify.Event) {
	// 	zap.S().Info("config file change", e.Name)
	// 	_ = v.ReadInConfig()
	// 	_ = v.Unmarshal(&global.ServerConfig)
	// 	zap.S().Info(global.ServerConfig)
	// })
}
