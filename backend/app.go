package backend

import (
	"bytes"
	"context"
	"encoding/base64"
	"image/png"

	sr "github.com/kbinani/screenshot"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) Init() {

}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.Init()
}

func (a *App) Shutdown(ctx context.Context) {

}

func (a *App) DomReady(ctx context.Context) {

}
func (a *App) Screenshot() (string, error) {
	// 获取所有屏幕的矩形区域
	n := sr.NumActiveDisplays()
	if n <= 0 {
		return "", nil
	}

	// 选择主屏幕进行截图
	img, err := sr.CaptureDisplay(0)
	if err != nil {
		return "", err
	}

	// 将图片编码为 PNG 格式
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return "", err
	}

	// 将图片字节数据编码为 Base64
	base64String := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64String, nil
}
