//go:build linux || darwin

package util

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type Util struct{}

func NewUtil() *Util {
	return &Util{}
}

// Db 获取数据库操作对象和数据库初始化
func (u *Util) Db() *gorm.DB {
	//打开数据库，如果不存在，则创建
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/data.db", u.GetAppPath())), &gorm.Config{})
	if err != nil {
		panic(any("数据库连接失败"))
	}
	return db
}

// Log 增加日志记录
func (u *Util) Log(content string) *Util {
	path := u.PathExist(fmt.Sprintf("%s/logs", u.GetAppPath()))
	// 创建或打开日志文件
	logFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", path, time.Now().Format("2006-01-02")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	//defer logFile.Close()
	//记录文件路径和行号
	_, file, line, _ := runtime.Caller(1)
	// 初始化日志
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(fmt.Sprintf("\n文件路径：%s:%d\n日志内容：%s\n", file, line, content))
	return u
}

// Schema 根据model自动建立数据表
func (u *Util) Schema(dst ...interface{}) {
	_ = u.Db().AutoMigrate(dst...)
}

// GetAppPath 获取应用主目录
func (u *Util) GetAppPath() string {
	//获取系统我的文档目录
	homeDir, dirErr := os.UserHomeDir()
	if dirErr != nil {
		panic(any("获取系统用户主目录失败"))
	}
	//获取我的文档目录
	return u.PathExist(fmt.Sprintf("%s/GoDeskFiles", homeDir))
}

// PathExist 判断文件目录是否存在，不存在创建
func (u *Util) PathExist(path string) string {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	return path
}

// ShellCMD 以shell方式运行cmd命令
func (u *Util) ShellCMD(cmdStr string, paramStr string) {
	exec.Command(cmdStr, paramStr).Start()
}

// SetSystemBackground 设置系统壁纸
func (u *Util) SetSystemBackground(path string) {
	if runtime.GOOS == "darwin" {
		//使用osascript执行AppleScript
		_ = exec.Command("osascript", "-e", fmt.Sprintf(`
		tell application "Finder"
			set desktop picture to POSIX file "%s"
		end tell
	`, path)).Start()
		//defaults命令不同版本参数不同
		//_ = exec.Command("defaults", "write", "com.apple.desktop", "picture", path).Start()
		return
	}
	// 设置壁纸
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", "file://"+path)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error setting wallpaper: %s", err)
	}
}

// HttpGet get请求
func (u *Util) HttpGet(url string) []byte {
	res, _ := http.Get(url)
	defer res.Body.Close()
	bytes, _ := io.ReadAll(res.Body)
	return bytes
}

// OpenDir 打开目录
func (u *Util) OpenDir(path string) {
	if runtime.GOOS == "darwin" {
		u.ShellCMD("open", path)
		return
	}
	if !regexp.MustCompile(`^(http|ftp)s?://`).MatchString(path) {
		path = fmt.Sprintf("file://%s", path)
	}
	u.ShellCMD("xdg-open", path)
}

// PathConvert 路径转换
func (u *Util) PathConvert(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}
