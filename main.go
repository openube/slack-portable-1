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
	Papp.ID = "slack-portable"
	Papp.Name = "Slack"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))
	roamingPath := CreateFolder(PathJoin(Papp.DataPath, "AppData", "Roaming", "Slack"))
	Log.Infof("Roaming path: %s", roamingPath)

	Papp.Process = PathJoin(Papp.AppPath, "Slack.exe")
	Papp.Args = nil
	Papp.WorkingDir = electronBinPath

	// Downloads folder
	downloadsPath := CreateFolder(PathJoin(Papp.Path, "downloads"))

	// Update slack settings
	Log.Info("Update Slack settings...")
	slackSettingsPath := PathJoin(roamingPath, "storage", "slack-settings")
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

	OverrideEnv("USERPROFILE", Papp.DataPath)
	Launch()
}
