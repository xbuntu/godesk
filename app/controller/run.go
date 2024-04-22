package controller

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"godesk/app/model"
)

// WailsRun 初始化
func WailsRun(assets embed.FS, port int, icon []byte) {
	// 创建控制器实例
	system := NewSystem()
	config := NewConfig()
	server := NewServer()
	wallpaper := NewWallpaper()
	audio := NewAudio()
	photo := NewPhoto()

	// 启动wails服务
	err := wails.Run(&options.App{
		Title:  "GoDesk：演示wails的使用",
		Width:  900,
		Height: 620,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        false,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			//设置 context 对象
			system.setCtx(ctx)
			config.setCtx(ctx)
			server.setCtx(ctx)
			audio.setCtx(ctx)
			photo.setCtx(ctx)
			wallpaper.setCtx(ctx)

			//启动自定义服务，初始化数据表
			server.start(port).log("程序已启动").schema(&model.Config{})
		},
		Bind: []interface{}{
			system,
			config,
			server,
			wallpaper,
			audio,
			photo,
		},
		Mac: &mac.Options{
			TitleBar: mac.TitleBarDefault(),
			About: &mac.AboutInfo{
				Title:   "GoDesk",
				Message: "wails2演示程序",
				Icon:    icon,
			},
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			DisableFramelessWindowDecorations: true,
		},
		Linux: &linux.Options{
			ProgramName:         "GoDesk",
			Icon:                icon,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyOnDemand,
			WindowIsTranslucent: false,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
