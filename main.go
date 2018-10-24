//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Slack.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	_ "github.com/kevinburke/go-bindata"
	"github.com/portapps/slack-portable/assets"

	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "slack-portable"
	Papp.Name = "Slack"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = CreateFolder(AppPathJoin("data"))

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))

	Papp.Process = PathJoin(Papp.AppPath, "Slack.exe")
	Papp.Args = nil
	Papp.WorkingDir = electronBinPath

	// Downloads folder
	downloadsPath := CreateFolder(PathJoin(Papp.Path, "downloads"))

	// Update slack settings
	Log.Info("Update Slack settings...")
	slackSettingsPath := PathJoin(Papp.DataPath, "storage", "slack-settings")
	if _, err := os.Stat(slackSettingsPath); err == nil {
		rawSettings, err := ioutil.ReadFile(slackSettingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			json.Unmarshal(rawSettings, &jsonMapSettings)
			Log.Info("Current settings:", jsonMapSettings)

			jsonMapSettings["resourcePath"] = PathJoin(electronBinPath, "resources", "app.asar")
			jsonMapSettings["PrefSSBFileDownloadPath"] = downloadsPath
			jsonMapSettings["notificationMethod"] = "html"
			Log.Info("New settings:", jsonMapSettings)

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				Log.Error("Slack settings marshal:", err)
			}
			err = ioutil.WriteFile(slackSettingsPath, jsonSettings, 0644)
			if err != nil {
				Log.Error("Write Slack settings:", err)
			}
		}
	} else {
		Log.Errorf("Slack settings not found in %s", slackSettingsPath)
	}

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Slack Portable.lnk")
	defaultShortcut, err := assets.Asset("Slack.lnk")
	if err != nil {
		Log.Error("Cannot load asset Discord.lnk:", err)
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error("Cannot write default shortcut:", err)
	}

	// Update default shortcut
	err = CreateShortcut(WindowsShortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       Papp.Process,
		Arguments:        WindowsShortcutProperty{Clear: true},
		Description:      WindowsShortcutProperty{Value: "Slack Portable by Portapps"},
		IconLocation:     WindowsShortcutProperty{Value: Papp.Process},
		WorkingDirectory: WindowsShortcutProperty{Value: Papp.AppPath},
	})
	if err != nil {
		Log.Error("Cannot create shortcut:", err)
	}

	Launch(os.Args[1:])
}
