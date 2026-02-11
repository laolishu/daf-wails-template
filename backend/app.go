/*
 * @Descripttion:
 * @version:
 * @Author: lfzxs@qq.com
 * @Date: 2026-02-05 15:38:06
 * @LastEditors: lfzxs@qq.com
 * @LastEditTime: 2026-02-10 16:38:03
 */
package backend

import (
	"context"
	"daf-wails-template/core/config"
	"daf-wails-template/core/logger"
	"daf-wails-template/core/systeminfo"
	"fmt"
	"io/fs"
	"log/slog"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// App struct
type App struct {
	ctx    context.Context
	config *config.Manager
	assets fs.FS
}

type LogWriteResult struct {
	Ok     bool   `json:"ok"`
	Error  string `json:"error,omitempty"`
	LogDir string `json:"logDir,omitempty"`
}

type ConfigSummary struct {
	Language string `json:"language"`
	LogLevel string `json:"logLevel"`
	LogDir   string `json:"logDir"`
}

// NewApp creates a new App application struct
func NewApp(assets fs.FS) *App {
	return &App{assets: assets}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetSystemInfo() systeminfo.Info {
	return systeminfo.Get()
}

func (a *App) GetWindowTitle() string {
	if a.config == nil {
		return config.DefaultWindowTitle
	}

	return a.config.GetWindowTitle()
}

func (a *App) GetLanguage() string {
	if a.config == nil {
		return config.DefaultLanguage
	}

	return a.config.GetLanguage()
}

func (a *App) SetLanguage(language string) error {
	if a.config == nil {
		return fmt.Errorf("config manager is not initialized")
	}

	switch language {
	case "zh-CN", "en-US":
		// allowed
	default:
		return fmt.Errorf("unsupported language: %s", language)
	}

	a.config.SetLanguage(language)
	return a.config.Save()
}

func (a *App) WriteTestLog() LogWriteResult {
	if a.config == nil {
		return LogWriteResult{Ok: false, Error: "config manager is not initialized"}
	}

	timestamp := time.Now().Format(time.RFC3339)
	message := fmt.Sprintf("[HealthCheck] log write test at %s", timestamp)
	slog.Info(message)

	return LogWriteResult{Ok: true, LogDir: a.config.GetLogDir()}
}

func (a *App) GetConfigSummary() (ConfigSummary, error) {
	if a.config == nil {
		return ConfigSummary{}, fmt.Errorf("config manager is not initialized")
	}

	return ConfigSummary{
		Language: a.config.GetLanguage(),
		LogLevel: a.config.GetLogLevel(),
		LogDir:   a.config.GetLogDir(),
	}, nil
}

func (a *App) initConfig() error {
	cfg, err := config.NewManager()
	if err != nil {
		return err
	}

	a.config = cfg

	return a.config.Load()
}

func (a *App) initLogger() error {
	if a.config == nil {
		return fmt.Errorf("config manager is not initialized")
	}

	_, err := logger.Init(logger.Config{
		Dir:            a.config.GetLogDir(),
		Level:          a.config.GetLogLevel(),
		RetentionDays:  a.config.GetLogRetentionDays(),
		ConsoleEnabled: true,
	})
	return err
}

func (a *App) Run() error {

	err := a.initConfig()
	if err != nil {
		return err
	}

	if err := a.initLogger(); err != nil {
		return err
	}

	err = wails.Run(&options.App{
		Title:           a.config.GetWindowTitle(),
		Width:           a.config.GetWindowWidth(),
		Height:          a.config.GetWindowHeight(),
		Frameless:       true,
		MinWidth:        config.DefaultWindowMinWidth,
		MinHeight:       config.DefaultWindowMinHeight,
		CSSDragProperty: "--wails-draggable",
		CSSDragValue:    "drag",
		AssetServer: &assetserver.Options{
			Assets: a.assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Logger:           logger.NewWailsAdapter(slog.Default()),
		OnStartup:        a.startup,
		Bind: []interface{}{
			a,
		},
	})

	slog.Info("启动成功")

	return err

}
