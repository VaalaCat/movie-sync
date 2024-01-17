package video

import (
	"movie-sync-server/conf"

	"github.com/iawia002/lux/extractors"
	_ "github.com/iawia002/lux/extractors/acfun"
	_ "github.com/iawia002/lux/extractors/bcy"
	_ "github.com/iawia002/lux/extractors/bilibili"
	_ "github.com/iawia002/lux/extractors/douyin"
	_ "github.com/iawia002/lux/extractors/douyu"
	_ "github.com/iawia002/lux/extractors/eporner"
	_ "github.com/iawia002/lux/extractors/facebook"
	_ "github.com/iawia002/lux/extractors/geekbang"
	_ "github.com/iawia002/lux/extractors/haokan"
	_ "github.com/iawia002/lux/extractors/hupu"
	_ "github.com/iawia002/lux/extractors/huya"
	_ "github.com/iawia002/lux/extractors/instagram"
	_ "github.com/iawia002/lux/extractors/iqiyi"
	_ "github.com/iawia002/lux/extractors/ixigua"
	_ "github.com/iawia002/lux/extractors/kuaishou"
	_ "github.com/iawia002/lux/extractors/mgtv"
	_ "github.com/iawia002/lux/extractors/miaopai"
	_ "github.com/iawia002/lux/extractors/netease"
	_ "github.com/iawia002/lux/extractors/pixivision"
	_ "github.com/iawia002/lux/extractors/pornhub"
	_ "github.com/iawia002/lux/extractors/qq"
	_ "github.com/iawia002/lux/extractors/reddit"
	_ "github.com/iawia002/lux/extractors/streamtape"
	_ "github.com/iawia002/lux/extractors/tangdou"
	_ "github.com/iawia002/lux/extractors/tiktok"
	_ "github.com/iawia002/lux/extractors/tumblr"
	_ "github.com/iawia002/lux/extractors/twitter"
	_ "github.com/iawia002/lux/extractors/udn"
	_ "github.com/iawia002/lux/extractors/universal"
	_ "github.com/iawia002/lux/extractors/vimeo"
	_ "github.com/iawia002/lux/extractors/vk"
	_ "github.com/iawia002/lux/extractors/weibo"
	_ "github.com/iawia002/lux/extractors/ximalaya"
	_ "github.com/iawia002/lux/extractors/xinpianchang"
	_ "github.com/iawia002/lux/extractors/xvideos"
	_ "github.com/iawia002/lux/extractors/yinyuetai"
	_ "github.com/iawia002/lux/extractors/youku"
	_ "github.com/iawia002/lux/extractors/youtube"
	_ "github.com/iawia002/lux/extractors/zhihu"
	"github.com/sirupsen/logrus"
)

func GetUrl(url string) string {
	data, err := extractors.Extract(url, extractors.Options{
		Playlist:         false,
		Items:            "",
		ItemStart:        0,
		ItemEnd:          0,
		ThreadNumber:     0,
		EpisodeTitleOnly: true,
		Cookie:           "",
		YoukuPassword:    "",
		YoukuCcode:       conf.ServerSetting.YoukuCcode,
		YoukuCkey:        conf.ServerSetting.YoukuCkey,
	})
	logrus.Info(data)
	if err != nil {
		logrus.Info(err)
		return ""
	}
	if len(data) == 0 {
		return ""
	}
	return data[0].URL
}
