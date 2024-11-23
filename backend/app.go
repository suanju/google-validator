package backend

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
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
	// 检查屏幕数量
	n := sr.NumActiveDisplays()
	if n <= 0 {
		return "", fmt.Errorf("截图出现错误")
	}

	// 捕获主屏幕截图
	img, err := sr.CaptureDisplay(0)
	if err != nil {
		return "", err
	}

	// 将图片编码为 PNG 并转为 Base64
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		return "", fmt.Errorf("error encoding image: %v", err)
	}

	base64Image := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return base64Image, nil
}
