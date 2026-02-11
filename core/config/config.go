package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"daf-wails-template/core/sysconfig"

	"github.com/spf13/viper"
)

const (
	keyLogDir           = "log.dir"
	keyLogLevel         = "log.level"
	keyLogRetentionDays = "log.retention_days"
	keyLanguage         = "i18n.language"
	keyWindowWidth      = "window.width"
	keyWindowHeight     = "window.height"
	keyWindowTitle      = "window.title"
)

type Manager struct {
	v          *viper.Viper
	configPath string
}

func NewManager() (*Manager, error) {
	configPath, err := resolveConfigPath()
	if err != nil {
		return nil, err
	}

	v := viper.New()
	v.SetDefault(keyLogDir, DefaultLogDir)
	v.SetDefault(keyLogLevel, DefaultLogLevel)
	v.SetDefault(keyLogRetentionDays, DefaultLogRetentionDays)
	v.SetDefault(keyLanguage, DefaultLanguage)
	v.SetDefault(keyWindowWidth, DefaultWindowWidth)
	v.SetDefault(keyWindowHeight, DefaultWindowHeight)
	v.SetDefault(keyWindowTitle, DefaultWindowTitle)
	v.SetConfigType("yaml")
	v.SetConfigFile(configPath)

	return &Manager{
		v:          v,
		configPath: configPath,
	}, nil
}

func (m *Manager) Load() error {
	if err := m.v.ReadInConfig(); err != nil {
		if isConfigNotFound(err) {
			if err := ensureConfigDir(m.configPath); err != nil {
				return err
			}
			return m.v.WriteConfigAs(m.configPath)
		}
		return err
	}
	return nil
}

func (m *Manager) Save() error {
	if err := ensureConfigDir(m.configPath); err != nil {
		return err
	}

	if _, err := os.Stat(m.configPath); err == nil {
		return m.v.WriteConfig()
	} else if os.IsNotExist(err) {
		return m.v.WriteConfigAs(m.configPath)
	} else {
		return err
	}
}

func (m *Manager) GetLogDir() string {
	return m.v.GetString(keyLogDir)
}

func (m *Manager) GetLogLevel() string {
	return m.v.GetString(keyLogLevel)
}

func (m *Manager) GetLogRetentionDays() int {
	return m.v.GetInt(keyLogRetentionDays)
}

func (m *Manager) GetLanguage() string {
	return m.v.GetString(keyLanguage)
}

func (m *Manager) SetLogDir(dir string) {
	m.v.Set(keyLogDir, dir)
}

func (m *Manager) SetLogLevel(level string) {
	m.v.Set(keyLogLevel, level)
}

func (m *Manager) SetLogRetentionDays(days int) {
	m.v.Set(keyLogRetentionDays, days)
}

func (m *Manager) SetLanguage(lang string) {
	m.v.Set(keyLanguage, lang)
}

func (m *Manager) GetWindowWidth() int {
	return m.v.GetInt(keyWindowWidth)
}

func (m *Manager) GetWindowHeight() int {
	return m.v.GetInt(keyWindowHeight)
}

func (m *Manager) GetWindowTitle() string {
	return m.v.GetString(keyWindowTitle)
}

func (m *Manager) SetWindowWidth(width int) {
	m.v.Set(keyWindowWidth, width)
}

func (m *Manager) SetWindowHeight(height int) {
	m.v.Set(keyWindowHeight, height)
}

func (m *Manager) SetWindowTitle(title string) {
	m.v.Set(keyWindowTitle, title)
}

func resolveConfigPath() (string, error) {
	configFile := sysconfig.GetConfigFile()
	if configFile == "" {
		return "", fmt.Errorf("config file name is empty")
	}

	configDir := sysconfig.GetConfigDir()
	if configDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		configDir = cwd
	}

	return filepath.Join(configDir, configFile), nil
}

func ensureConfigDir(configPath string) error {
	dir := filepath.Dir(configPath)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

func isConfigNotFound(err error) bool {
	var notFound viper.ConfigFileNotFoundError
	if errors.As(err, &notFound) {
		return true
	}
	return os.IsNotExist(err)
}
