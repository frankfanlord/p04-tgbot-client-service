package webhook

import (
	"client-service/config"
	"io"
	"jarvis/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PathWebhook          = "webhook"
	CtxHeaderSecretToken = "X-Telegram-Bot-Api-Secret-Token"
)

func handler(ctx *gin.Context) {
	defer func() { _ = ctx.Request.Body }()

	secret := ctx.GetHeader(CtxHeaderSecretToken)

	if secret != config.Instance().Telegram.SecretToken {
		logger.App().Warnf("receive wrong secret token : %s - %s", secret, config.Instance().Telegram.SecretToken)
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.App().Errorf("ReadAll from request body error : %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	logger.App().Infof("receive update : %s", string(data))

	go func(vs []byte) { _channel <- vs }(data[:])

	ctx.JSON(http.StatusOK, nil)
}
