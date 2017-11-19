//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	. "github.com/portapps/portapps"
)

func init() {
	App.ID = "slack-portable"
	App.Name = "Slack"
	Init()
}

func main() {
	App.MainPath = FindElectronMainFolder("app-")
	App.DataPath = CreateFolder(PathJoin(App.RootDataPath, "AppData", "Roaming", "Slack"))
	App.Process = RootPathJoin("Slack.exe")
	App.Args = nil
	App.WorkingDir = App.MainPath

	// Downloads folder
	downloadsPath := CreateFolder(PathJoin(App.Path, "downloads"))

	// Update slack settings
	Log.Info("Update Slack settings...")
	slackSettingsPath := PathJoin(App.DataPath, "storage", "slack-settings")
	if _, err := os.Stat(slackSettingsPath); err == nil {
		rawSettings, err := ioutil.ReadFile(slackSettingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			json.Unmarshal(rawSettings, &jsonMapSettings)
			Log.Info("Current settings:", jsonMapSettings)

			jsonMapSettings["resourcePath"] = PathJoin(App.MainPath, "resources", "app.asar")
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

	OverrideEnv("USERPROFILE", App.RootDataPath)
	Launch()
}
