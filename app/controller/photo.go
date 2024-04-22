package controller

import (
	"fmt"
	"strings"
)

// Photo 图片
type Photo struct {
	Base
}

func NewPhoto() *Photo {
	return &Photo{}
}

// PhotoItem 全景图配置
type PhotoItem struct {
	Id        string            `json:"id"`
	Panorama  interface{}       `json:"panorama"`
	Thumbnail interface{}       `json:"thumbnail"`
	Name      string            `json:"name"`
	Options   map[string]string `json:"options"`
}

// GetList 获取列表
func (p Photo) GetList() []PhotoItem {
	image := `https://bbsstatic.oneplus.com/data/attachment/forum/202012/08/055151weo5cowlcludb32j.jpg
https://bbsstatic.oneplus.com/data/attachment/forum/202012/08/055151weo5cowlcludb32j.jpg
https://bbsstatic.oneplus.com/data/attachment/forum/202012/08/055151weo5cowlcludb32j.jpg
https://bbsstatic.oneplus.com/data/attachment/forum/202012/08/055151weo5cowlcludb32j.jpg`
	list := make([]PhotoItem, 0)
	images := strings.Split(image, "\n")
	for id, url := range images {
		list = append(list, PhotoItem{
			Id:        fmt.Sprintf("id%d", id),
			Panorama:  url,
			Thumbnail: url,
			Name:      fmt.Sprintf("测试%d", id),
			Options: map[string]string{
				"caption": fmt.Sprintf("测试%d", id),
			},
		})
	}
	return list
}
