package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"godesk/app/model"
	"godesk/app/util"
	"strings"
)

type Wallpaper struct{}

func NewWallpaper() *Wallpaper {
	return &Wallpaper{}
}

// BingRes 必应壁纸每日返回结构
type BingRes struct {
	LastUpdate string `json:"LastUpdate"`
	Total      int    `json:"Total"`
	Language   string `json:"Language"`
	Message    string `json:"message"`
	Status     bool   `json:"status"`
	Success    bool   `json:"success"`
	Info       string `json:"info"`
	Data       []struct {
		Startdate     string        `json:"startdate"`
		Fullstartdate string        `json:"fullstartdate"`
		Enddate       string        `json:"enddate"`
		Url           string        `json:"url"`
		Urlbase       string        `json:"urlbase"`
		Copyright     string        `json:"copyright"`
		Copyrightlink string        `json:"copyrightlink"`
		Title         string        `json:"title"`
		Quiz          string        `json:"quiz"`
		Wp            bool          `json:"wp"`
		Hsh           string        `json:"hsh"`
		Drk           int           `json:"drk"`
		Top           int           `json:"top"`
		Bot           int           `json:"bot"`
		Hs            []interface{} `json:"hs"`
	} `json:"data"`
}

// GetCate 获取分类
func (w *Wallpaper) GetCate() []model.WallpaperCate {
	//初始化20个分类
	cateList := make([]model.WallpaperCate, 0)
	cateString := "风景,美女,爱情,卡通,明星,动物,游戏,汽车,影视,体育"
	split := strings.Split(cateString, ",")
	for i, title := range split {
		cateList = append(cateList, model.WallpaperCate{
			ID:    i + 1,
			Title: title,
		})
	}
	return cateList
}

// GetList 获取壁纸
func (w *Wallpaper) GetList() ([]model.Wallpaper, int) {
	//初始化壁纸
	var res BingRes
	u := util.NewUtil()
	_ = json.NewDecoder(bytes.NewReader(u.HttpGet("https://raw.onmicrosoft.cn/Bing-Wallpaper-Action/main/data/zh-CN_all.json"))).Decode(&res)
	list := make([]model.Wallpaper, 0)
	for i, data := range res.Data {
		list = append(list, model.Wallpaper{
			CID:   i,
			Title: data.Title,
			Tag:   data.Title,
			Url:   fmt.Sprintf("https://s.cn.bing.net%s", data.Url),
		})
	}
	return list, res.Total
}
