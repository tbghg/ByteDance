package service

import (
	"ByteDance/cmd/video/repository"
	"time"
)

func GetVideoFeed(lastTime int64) {

	formatLastTime := time.Unix(lastTime, 0).Format("2006-01-02 15:04:05")
	repository.VideoDao.GetVideoFeed(formatLastTime)
	// 去数据库里面找时间小于lastTimed的数据，找10个回来，每一个填充作者信息，join起来
	// videoList :=
	// 把最后一个视频的time作为nextTime返回

	// return NextTime和videoList

}
