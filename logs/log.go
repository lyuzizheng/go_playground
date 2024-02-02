package logs

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func init() {
	base, _ := zap.NewDevelopment()
	Logger = base.Sugar()
}
