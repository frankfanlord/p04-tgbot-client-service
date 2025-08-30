package telegram

import (
	"client-service/core/telegram/basic"
	"jarvis/logger"
	"jarvis/middleware/mq/nats"
)

// Init 初始化
func Init(domain, token, webhook, secretToken string, channel <-chan []byte) error {
	logger.App().Infoln("=================================================== start init telegram bot ===================================================")
	defer logger.App().Infoln("=================================================== stop init telegram bot ===================================================")

	if err := basic.Init(domain, token); err != nil {
		return err
	}

	_bot = &bot{
		webhook:     webhook,
		secretToken: secretToken,
		channel:     channel,
		close:       make(chan struct{}),
		done:        make(chan struct{}),
	}

	if err := _bot.Initialize(); err != nil {
		return err
	}

	// ============= nats
	if subscription, err := nats.Instance().QueueSubscribe(JSPinSubject, JSClientQueue, doPin); err != nil {
		return err
	} else {
		_subscriptions = append(_subscriptions, subscription)
		logger.App().Infof("=========== subscribe to [%s][%s] success ===========", JSPinSubject, JSClientQueue)
	}

	if subscription, err := nats.Instance().QueueSubscribe(JSDeleteSubject, JSClientQueue, doPin); err != nil {
		return err
	} else {
		_subscriptions = append(_subscriptions, subscription)
		logger.App().Infof("=========== subscribe to [%s][%s] success ===========", JSDeleteSubject, JSClientQueue)
	}

	// ===================
	if subscription, err := nats.Instance().QueueSubscribe(SSMissionResponseSubject, CSQueue, doResponse); err != nil {
		return err
	} else {
		_subscriptions = append(_subscriptions, subscription)
		logger.App().Infof("=========== subscribe to [%s][%s] success ===========", SSMissionResponseSubject, CSQueue)
	}

	return nil
}

// Start 启动
func Start(channel chan<- any) {
	logger.App().Infoln("=================================================== start telegram bot ===================================================")
	defer logger.App().Infoln("=================================================== stop telegram bot ===================================================")

	channel <- nil

	_bot.Run()
}

// Shutdown 关闭
func Shutdown() error {
	logger.App().Infoln("=================================================== start shutdown telegram bot ===================================================")
	defer logger.App().Infoln("=================================================== stop shutdown telegram bot ===================================================")

	if _subscriptions != nil && len(_subscriptions) > 0 {
		for idx := range _subscriptions {
			subscription := _subscriptions[idx]
			_ = subscription.Unsubscribe()
		}
	}

	_bot.Shutdown()

	return nil
}

// LoadCache 加载缓存
func LoadCache() error {
	logger.App().Infoln("=================================================== telegram bot start load cache ===================================================")
	defer logger.App().Infoln("=================================================== telegram bot stop load cachek ===================================================")
	return nil
}
