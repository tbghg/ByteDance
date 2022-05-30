package repository

import (
	"sync"
)

type VideoDaoStruct struct {
}

var (
	VideoDao  *VideoDaoStruct
	videoOnce sync.Once
)

func init() {
	videoOnce.Do(func() {
		VideoDao = &VideoDaoStruct{}
	})
}

func (*VideoDaoStruct) GetVideoFeed(lastTime string) {
	//v := dal.ConnQuery.Video
	//u := dal.ConnQuery.User
	// 内联查询
	//err := v.Select(u.ID, u.Username, v.ID, v.PlayURL, v.CoverURL, v.Time, v.Title).Join(u, u.ID.EqCol(v.AuthorID)).Scan()

}
