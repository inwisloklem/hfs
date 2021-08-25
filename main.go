package main

import (
	_ "embed"
	"hfs/core"
	"hfs/settings"

	"github.com/wailsapp/wails"
)

const configFile = "config.toml"

var (
	//go:embed frontend/build/static/css/main.css
	css string
	//go:embed frontend/build/static/js/main.js
	js string
)

func startApp(c *core.Control) {
	app := wails.CreateApp(&wails.AppConfig{
		CSS:    css,
		JS:     js,
		Title:  "HFS",
		Height: 124,
		Width:  496,
	})
	app.Bind(c)
	app.Run()
}

func main() {
	config, err := settings.ReadConfigFile(configFile)
	control := &core.Control{
		Config:      config,
		HasNoConfig: false,
		Message:     "Click 'Save' to save Profile...",
	}
	if err != nil {
		control.SetErrorMessage(err)
		control.HasNoConfig = true
		startApp(control)
		return
	}
	control.Message = "Click 'Save' to save " + control.GetProfileName()
	startApp(control)
}
