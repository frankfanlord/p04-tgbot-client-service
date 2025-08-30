package core

import (
	"client-service/core/telegram"
	"client-service/core/webhook"
	"errors"
	"jarvis/logger"
)

var (
	channel = make(chan []byte, 10000)
)

func Init(prefix, addr, domain, token, wbURL, secretToken string) error {
	logger.App().Infoln("=================================================== start init core ===================================================")
	defer logger.App().Infoln("=================================================== stop init core ===================================================")

	if err := webhook.Init(prefix, addr, channel); err != nil {
		return err
	}

	if err := telegram.Init(domain, token, wbURL, secretToken, channel); err != nil {
		return err
	}

	return nil
}

func Start() error {
	logger.App().Infoln("=================================================== start core ===================================================")
	defer logger.App().Infoln("=================================================== stop core ===================================================")

	channel := make(chan any, 2)

	go webhook.Start(channel)
	go telegram.Start(channel)

	count := 2
	var err error

	for i := count; i > 0; i-- {
		v, ok := <-channel
		if !ok {
			err = errors.New("channel being closed")
			break
		}

		if v == nil {
			continue
		}

		if ve, yes := v.(error); yes {
			err = ve
			break
		}
	}

	return err
}

func Shutdown() error {
	logger.App().Infoln("=================================================== start shutdown core ===================================================")
	defer logger.App().Infoln("=================================================== stop shutdown core ===================================================")

	if err := webhook.Shutdown(); err != nil {
		return err
	}

	if err := telegram.Shutdown(); err != nil {
		return err
	}

	return nil
}

func LoadCache() error {
	logger.App().Infoln("=================================================== start load core cache ===================================================")
	defer logger.App().Infoln("=================================================== stop load core cache ===================================================")

	if err := webhook.LoadCache(); err != nil {
		return err
	}

	if err := telegram.LoadCache(); err != nil {
		return err
	}

	return nil
}
