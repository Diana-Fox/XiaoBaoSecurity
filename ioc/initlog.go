package ioc

import "go.uber.org/zap"

func InitLog() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	//替换全局的对象
	zap.ReplaceGlobals(logger)
}
