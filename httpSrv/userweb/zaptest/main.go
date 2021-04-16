package main

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig() //配置输出到文件
	cfg.OutputPaths = []string{
		"./myproject.log", //两者都输出
		"stdout",
	}
	return cfg.Build()
}

func main() {
	// logger, _ := zap.NewProduction()
	//logger, _ := zap.NewDevelopment()
	logger, _ := NewLogger()
	defer logger.Sync() //flush buffer , if any
	url := "https://imooc.com"
	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch URL", "url", url, "attempt", 3)
	// sugar.Infof("failed to fetch URL:%s", url)  会使用反射机制，相对较慢
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
	)
}
