//go:build windows

package util

import (
	"fmt"
	"golang.org/x/sys/windows"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

var (
	moduser32                 = windows.NewLazySystemDLL("user32.dll")
	procSystemParametersInfoW = moduser32.NewProc("SystemParametersInfoW")
)

type Util struct{}

func NewUtil() *Util {
	return &Util{}
}

// Db 获取数据库操作对象和数据库初始化
func (u *Util) Db() *gorm.DB {
	//打开数据库，如果不存在，则创建
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s\\data.db", u.GetAppPath())), &gorm.Config{})
	if err != nil {
		panic(any("数据库连接失败"))
	}
	return db
}

// Log 增加日志记录
func (u *Util) Log(content string) *Util {
	path := u.PathExist(fmt.Sprintf("%s\\logs", u.GetAppPath()))
	// 创建或打开日志文件
	logFile, err := os.OpenFile(fmt.Sprintf("%s\\%s.log", path, time.Now().Format("2006-01-02")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
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
	return u.PathExist(fmt.Sprintf("%s\\Documents\\GoDesk Files", homeDir))
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
	open, _ := syscall.UTF16PtrFromString("open")
	cmd, _ := syscall.UTF16PtrFromString(cmdStr)
	param, _ := syscall.UTF16PtrFromString(paramStr)
	_ = windows.ShellExecute(0, open, cmd, param, nil, windows.SW_SHOWNORMAL)
}

// SetSystemBackground 设置系统壁纸
func (u *Util) SetSystemBackground(path string) {
	filename, _ := syscall.UTF16PtrFromString(path)
	_, _, _ = syscall.SyscallN(procSystemParametersInfoW.Addr(), 20, 1, uintptr(unsafe.Pointer(filename)), 0x1|0x2)
}

// HttpGet get请求
func (u *Util) HttpGet(url string) []byte {
	res, _ := http.Get(url)
	defer res.Body.Close()
	bytes, _ := io.ReadAll(res.Body)
	return bytes
}

// OpenDir 打开指定目录
func (u *Util) OpenDir(path string) {
	u.ShellCMD("explorer", path)
}

// PathConvert 路径转换
func (u *Util) PathConvert(path string) string {
	return path
}
