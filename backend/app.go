package backend

import (
	"context"
	"fmt"

	sr "github.com/kbinani/screenshot"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
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

// 二维码功能
func (a *App) Screenshot() (string, error) {
	code := ""
	// 获取所有屏幕的矩形区域
	n := sr.NumActiveDisplays()
	if n <= 0 {
		return code, fmt.Errorf("no active displays found")
	}

	// 选择主屏幕进行截图
	img, err := sr.CaptureDisplay(0)
	if err != nil {
		return code, err
	}

	if err != nil {
		fmt.Println("Error decoding image:", err)
		return code, err
	}

	// 将图片转换为 BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		fmt.Println("Error creating binary bitmap:", err)
		return code, err
	}

	// 使用二维码解码器
	reader := qrcode.NewQRCodeReader()
	result, err := reader.Decode(bmp, nil)
	if err != nil {
		fmt.Println("Error decoding QR code:", err)
		return code, err
	}
	code = result.String()

	return code, nil
}
