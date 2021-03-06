// +build windows

package app

import (
	"os"
	"path/filepath"
)

func (app *App) osInit() {
	home := os.Getenv("USERPROFILE")
	if len(home) < 1 {
		home = `C:\Users\Administrator`
	}
	localappdata := os.Getenv("LOCALAPPDATA")
	if len(localappdata) < 1 {
		localappdata = os.Getenv("APPDATA")
		if len(localappdata) < 1 {
			localappdata = filepath.Join(home, ".local", "share")
		}
	}
	programData := os.Getenv("ProgramData")
	if len(programData) < 1 {
		programData = os.Getenv("ALLUSERSPROFILE")
		if len(programData) < 1 {
			programData = `C:\ProgramData`
		}
	}
	xdg_cache_home := os.Getenv("XDG_CACHE_HOME")
	if len(xdg_cache_home) < 1 {
		xdg_cache_home = filepath.Join(localappdata, "Microsoft", "Windows", "Temporary Internet Files")
	}
	xdg_config_home := os.Getenv("XDG_CONFIG_HOME")
	if len(xdg_config_home) < 1 {
		xdg_config_home = localappdata
	}
	xdg_data_home := os.Getenv("XDG_DATA_HOME")
	if len(xdg_data_home) < 1 {
		xdg_data_home = localappdata
	}

	app.cache = filepath.Join(xdg_cache_home, app.path)
	app.data = filepath.Join(xdg_data_home, app.path)
	app.desktop = filepath.Join(home, "Desktop")
	app.documents = filepath.Join(home, "Documents")
	app.downloads = filepath.Join(home, "Downloads")
	app.home = home
	app.localConfig = filepath.Join(xdg_config_home, app.path)
	app.man = filepath.Join(programData, "Common", "man")
	app.pictures = filepath.Join(home, "Pictures")
	app.screenshots = filepath.Join(app.pictures, "Screenshots")
	app.systemConfig = filepath.Join(programData, app.path)
}
