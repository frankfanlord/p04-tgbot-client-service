package telegram

import (
	"client-service/config"
	"client-service/core/telegram/basic"
	"client-service/core/telegram/basic/structure"
	"context"
	"errors"
	"jarvis/dao/db/redis"
	"jarvis/logger"
	"net/url"
	"path/filepath"

	"github.com/duke-git/lancet/v2/fileutil"
	ORedis "github.com/redis/go-redis/v9"
)

var (
	ErrEmptyWebhook error = errors.New("empty webhook")
	ErrEmptyScheme  error = errors.New("empty scheme")
	ErrEmptyHost    error = errors.New("empty host")
)

func (b *bot) Initialize() error {
	logger.App().Infof("================================== initialize start ==================================")
	defer logger.App().Infof("================================== initialize finish ==================================")

	if b.webhook == "" {
		return ErrEmptyWebhook
	}

	whURL, err := url.Parse(b.webhook)
	if err != nil {
		return err
	}
	if whURL.Scheme == "" {
		return ErrEmptyScheme
	}
	if whURL.Host == "" {
		return ErrEmptyHost
	}

	if err := initFiles(); err != nil {
		return err
	}

	// if err := initCommands(); err != nil {
	// 	return err
	// }
	// if err := initName("快搜🔍中文搜索@KSOU"); err != nil {
	// 	return err
	// }
	// if err := initDescription("🔍这是一个搜索引擎，向它发送关键词来寻找群组、频道、电影、音乐、视频"); err != nil {
	// 	return err
	// }
	// if err := initShortDescription("帮你找到有趣的群组、频道、视频、音乐、电影、新闻.官方公告:@Pabl02025Bot"); err != nil {
	// 	return err
	// }
	// if err := initWebhook(b.webhook, b.secretToken); err != nil {
	// 	return err
	// }
	//
	// me, gmErr := basic.GetMe()
	// if gmErr != nil {
	// 	return gmErr
	// }
	//
	// logger.App().Infof("Me : %+v", *(me))

	return nil
}

func initFiles() error {
	filesPath := filepath.Join(config.Instance().Runtime.WD, "files")

	if !fileutil.IsExist(filesPath) || !fileutil.IsDir(filesPath) {
		return errors.New("there is not files : " + filesPath)
	}

	for _, field := range []string{
		VideoCL, VideoFMV, VideoR18Desc, VideoBuildG, VideoStart,
	} {
		value := ""
		if cmd := redis.Instance().HGet(context.Background(), RKVideoFileID, field); cmd.Err() != nil {
			if cmd.Err() != ORedis.Nil {
				return cmd.Err()
			}
		} else {
			value = cmd.Val()
		}

		if value == "" {
			thePath := filepath.Join(filesPath, field)
			if !fileutil.IsExist(thePath) {
				logger.App().Errorf("[%s] not in redis and also not in files dir : [%s]", field, thePath)
				continue
			}

			if fileID, uvErr := basic.UploadVideo(6683769283, field, thePath); uvErr != nil {
				return uvErr
			} else {
				logger.App().Infof("upload [%s] success : %s", field, fileID)
				if cmd := redis.Instance().HSet(context.Background(), RKVideoFileID, field, fileID); cmd.Err() != nil {
					logger.App().Errorf("hset [%s]-[%s] error : [%s]", field, fileID, cmd.Err().Error())
				}
				_videoFileIDMap[field] = value
			}
		} else {
			_videoFileIDMap[field] = value
		}
	}

	return nil
}

func initCommands() error {
	languages := []string{"zh", "en", ""}

	var err error

	for _, lang := range languages {
		if err = basic.DeleteMyCommands("default", lang); err != nil {
			break
		}
	}

	if err != nil {
		return err
	}

	return basic.SetMyCommands(
		[]*structure.BotCommand{
			{Command: "/reso", Description: "🔍热搜排行榜"},
			{Command: "/daoh", Description: "👥群组导航"},
			{Command: "/help", Description: "❓教程帮助"},
			{Command: "/more", Description: "❄️更多功能"},
			{Command: "/privacy", Description: "📄隐私政策"},
		},
	)
}

func initName(name string) error {
	info, gErr := basic.GetMyName()
	if gErr != nil {
		return gErr
	}

	if info.Name == name {
		return nil
	}

	return basic.SetMyName(name)
}

func initDescription(desc string) error {
	info, gErr := basic.GetMyDescription()
	if gErr != nil {
		return gErr
	}

	if info.Description == desc {
		return nil
	}

	return basic.SetMyDescription(desc)
}

func initShortDescription(desc string) error {
	info, gErr := basic.GetMyShortDescription()
	if gErr != nil {
		return gErr
	}

	if info.Description == desc {
		return nil
	}

	return basic.SetMyShortDescription(desc)
}

func initWebhook(webhook, secretToken string) error {
	// 1.get webhook info
	info, gwiErr := basic.GetWebhookInfo()
	if gwiErr != nil {
		return gwiErr
	}

	// 2.if webhook info is setup already , compare with currently one
	if info.URL != "" && info.URL == webhook {
		return nil
	}

	if info.URL != "" && info.URL != webhook {
		if err := basic.DeleteWebhook(); err != nil {
			return err
		}
	}

	// 3.setup currently one
	return basic.SetWebhook(webhook, secretToken)
}

// ========================================================================================================

const (
	RKVideoFileID = "VideoFileIDMap"

	VideoCL      = "changeLanuage.mp4"
	VideoFMV     = "freeMusicDesc.mp4"
	VideoR18Desc = "release18Desc.mp4"
	VideoBuildG  = "buildGroupDesc.mp4"
	VideoStart   = "start.mp4"
)

var _videoFileIDMap = map[string]string{}
