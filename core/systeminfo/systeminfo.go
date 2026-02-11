package systeminfo

import (
	"runtime"

	"daf-wails-template/core/sysconfig"
)

type Info struct {
	Version     string `json:"version"`
	BuildTime   string `json:"buildTime"`
	Environment string `json:"environment"`
}

func Get() Info {
	info := sysconfig.GetInfo()
	return Info{
		Version:     info.Version,
		BuildTime:   info.BuildTime,
		Environment: runtime.GOOS + "/" + runtime.GOARCH,
	}
}
