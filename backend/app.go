package backend

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"

	"github.com/kbinani/screenshot"
	"github.com/liyue201/goqr"
	"github.com/nfnt/resize"
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

func (a *App) ScreenshotRecognize() ([]string, error) {
	// 检查屏幕数量
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		return nil, fmt.Errorf("没有可用屏幕进行截图")
	}

	// 捕获主屏幕截图
	img, err := screenshot.CaptureDisplay(0)
	if err != nil {
		return nil, fmt.Errorf("截图失败: %v", err)
	}

	// 将截图编码为 PNG 格式
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		return nil, fmt.Errorf("图片编码失败: %v", err)
	}

	// 解码图片
	decodedImg, _, err := image.Decode(bytes.NewReader(buffer.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("图片解码失败: %v", err)
	}

	// 可选：缩放图片，提升二维码识别效率
	decodedImg = resize.Resize(800, 0, decodedImg, resize.Lanczos3)

	// 识别二维码
	qrCodes, err := goqr.Recognize(decodedImg)
	if err != nil {
		return nil, fmt.Errorf("二维码识别失败: %v", err)
	}

	// 提取二维码内容
	var results []string
	for _, code := range qrCodes {
		results = append(results, string(code.Payload))
	}

	// 如果没有识别到二维码
	if len(results) == 0 {
		return nil, fmt.Errorf("未检测到二维码")
	}

	return results, nil
}
