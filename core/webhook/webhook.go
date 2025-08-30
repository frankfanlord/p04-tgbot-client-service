package webhook

import (
	"context"
	"errors"
	"jarvis/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	_server  *http.Server
	_channel chan<- []byte
)

// Init 初始化
func Init(prefix, addr string, channel chan<- []byte) error {
	logger.App().Infoln("=================================================== start init webhook ===================================================")
	defer logger.App().Infoln("=================================================== stop init webhook ===================================================")

	if channel == nil {
		return errors.New("channel is nil")
	}

	_channel = channel

	gin.DefaultWriter = logger.GinWriter(logrus.Fields{"component": "webhook"})
	gin.SetMode(gin.DebugMode)

	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())

	grouper := engine.Group(prefix)

	grouper.POST(PathWebhook, handler)

	_server = &http.Server{Addr: addr, Handler: engine}

	return nil
}

// Start 启动
func Start(channel chan<- any) {
	logger.App().Infoln("=================================================== start webhook ===================================================")
	defer logger.App().Infoln("=================================================== stop webhook ===================================================")

	if _server == nil {
		channel <- errors.New("server is nil")
		return
	}

	channel <- nil

	if err := _server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.App().Errorf("failed to start server: %s", err.Error())
	}
}

// Shutdown 关闭
func Shutdown() error {
	logger.App().Infoln("=================================================== start shutdown webhook ===================================================")
	defer logger.App().Infoln("=================================================== stop shutdown webhook ===================================================")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()
	return _server.Shutdown(ctx)
}

// LoadCache 加载缓存
func LoadCache() error {
	logger.App().Infoln("=================================================== webhook start load cache ===================================================")
	defer logger.App().Infoln("=================================================== webhook stop load cache ===================================================")
	return nil
}
