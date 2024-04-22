package controller

import (
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path"
	"strings"
)

// Audio 音频播放
type Audio struct {
	Base
}

func NewAudio() *Audio {
	return &Audio{}
}

// AudioRes 音频信息
type AudioRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Src  string `json:"src"`
	Lrc  string `json:"lrc"`
}

// GetList 获取目录下的音频文件
func (a *Audio) GetList() Res {
	list := make([]AudioRes, 0)
	dir, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if len(dir) == 0 {
		return a.error("请选择目录")
	}
	// 读取目录
	files, _ := os.ReadDir(dir)
	// 过滤音频文件
	index := 0
	for _, file := range files {
		if file.IsDir() {
			continue // 跳过子目录
		}
		ext := path.Ext(file.Name())
		if ext == ".mp3" || ext == ".m4a" {
			name := strings.TrimSuffix(file.Name(), ext)
			audio := AudioRes{
				Id:   index,
				Name: name,
				Path: a.pathConvert(fmt.Sprintf("%s\\%s", dir, file.Name())),
				Src:  a.pathConvert(fmt.Sprintf("http://localhost:%d/file?filename=%s\\%s", 10086, dir, file.Name())),
				Lrc:  "[00:00.00]歌词文件需要和歌曲在同一目录，且名称一样，扩展名为.lrc",
			}
			lrcFile := a.pathConvert(fmt.Sprintf("%s\\%s.lrc", dir, name))
			if _, err := os.Stat(lrcFile); !os.IsNotExist(err) {
				bytes, _ := os.ReadFile(lrcFile)
				audio.Lrc = string(bytes)
			}
			list = append(list, audio)
			index = index + 1
		}
	}
	a.SaveToJson(list)
	return a.success(list)
}

// SaveToJson 把播放列表存入json文件
func (a *Audio) SaveToJson(list []AudioRes) {
	file, _ := os.Create(a.pathConvert(fmt.Sprintf("%s\\audio.json", a.getAppPath())))
	defer file.Close()

	// 创建json编码器
	_ = json.NewEncoder(file).Encode(list)
}

// GetFromJson 从json文件中获取
func (a *Audio) GetFromJson() (list []AudioRes) {
	file := a.pathConvert(fmt.Sprintf("%s\\audio.json", a.getAppPath()))
	// 创建json编码器
	bytes, _ := os.ReadFile(file)
	_ = json.Unmarshal(bytes, &list)
	return
}
