package controller

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"godesk/app/service"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

type Wallpaper struct {
	Base
}

func NewWallpaper() *Wallpaper {
	return &Wallpaper{}
}

// SetBackGround 下载远程图片并设置成壁纸
func (w *Wallpaper) SetBackGround(url string) {
	//获取壁纸保存名称
	wPath := w.pathExist(w.pathConvert(fmt.Sprintf("%s\\wallpapers", w.getAppPath())))
	split := strings.Split(url, ".")
	ext := split[len(split)-1]
	filename := w.pathConvert(fmt.Sprintf("%s\\%d.%s", wPath, time.Now().Unix(), ext))

	//下载
	_ = os.WriteFile(filename, w.httpGet(url), os.ModePerm)

	//设置壁纸
	w.setSystemBackground(filename)
}

// SaveWallpaper 保存壁纸到指定目录
func (w Wallpaper) SaveWallpaper(fileUrl string) Res {
	f, _ := url.Parse(fileUrl)
	// 使用 path.Base 获取路径的基本名称
	fileName := path.Base(f.Path)
	wPath, err := runtime.SaveFileDialog(w.ctx, runtime.SaveDialogOptions{
		DefaultFilename: fileName,
		Title:           "保存当前壁纸",
	})
	if err != nil || len(wPath) < 1 {
		return w.error("已取消")
	}
	//下载
	_ = os.WriteFile(wPath, w.httpGet(fileUrl), os.ModePerm)
	return w.success(nil)
}

// GetCate 获取所有分类
func (w *Wallpaper) GetCate() Res {
	return w.success(service.NewWallpaper().GetCate())
}

// GetList 获取所有壁纸
func (w *Wallpaper) GetList() Res {
	list, total := service.NewWallpaper().GetList()
	res := make(map[string]interface{})
	res["list"] = list
	res["total"] = total
	return w.success(res)
}
