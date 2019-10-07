//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Slack.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/shortcut"
	"github.com/portapps/portapps/pkg/utl"
	"github.com/portapps/slack-portable/assets"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("slack-portable", "Slack"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	electronBinPath := utl.PathJoin(app.AppPath, utl.FindElectronAppFolder("app-", app.AppPath))

	app.Process = utl.PathJoin(electronBinPath, "Slack.exe")
	app.Args = []string{
		"--user-data-dir=" + app.DataPath,
	}
	app.WorkingDir = electronBinPath

	// Downloads folder
	downloadsPath := utl.CreateFolder(app.RootPath, "downloads")

	// Update slack settings
	Log.Info().Msg("Update Slack settings...")
	slackSettingsPath := utl.PathJoin(app.DataPath, "storage", "slack-settings")
	if _, err := os.Stat(slackSettingsPath); err == nil {
		rawSettings, err := ioutil.ReadFile(slackSettingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			if err = json.Unmarshal(rawSettings, &jsonMapSettings); err != nil {
				Log.Error().Err(err).Msg("Settings unmarshal")
			}
			Log.Info().Interface("settings", jsonMapSettings).Msg("Current settings")

			jsonMapSettings["resourcePath"] = utl.PathJoin(electronBinPath, "resources", "app.asar")
			jsonMapSettings["PrefSSBFileDownloadPath"] = downloadsPath
			jsonMapSettings["notificationMethod"] = "html"
			Log.Info().Interface("settings", jsonMapSettings).Msg("New settings")

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				Log.Error().Err(err).Msg("Settings marshal")
			}
			err = ioutil.WriteFile(slackSettingsPath, jsonSettings, 0644)
			if err != nil {
				Log.Error().Err(err).Msg("Write settings")
			}
		}
	} else {
		Log.Error().Msgf("Slack settings not found in %s", slackSettingsPath)
	}

	// Copy default shortcut
	shortcutPath := path.Join(utl.StartMenuPath(), "Slack Portable.lnk")
	defaultShortcut, err := assets.Asset("Slack.lnk")
	if err != nil {
		Log.Error().Err(err).Msg("Cannot load asset Slack.lnk")
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error().Err(err).Msg("Cannot write default shortcut")
	}

	// Update default shortcut
	err = shortcut.Create(shortcut.Shortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       app.Process,
		Arguments:        shortcut.Property{Clear: true},
		Description:      shortcut.Property{Value: "Slack Portable by Portapps"},
		IconLocation:     shortcut.Property{Value: app.Process},
		WorkingDirectory: shortcut.Property{Value: app.AppPath},
	})
	if err != nil {
		Log.Error().Err(err).Msg("Cannot create shortcut")
	}
	defer func() {
		if err := os.Remove(shortcutPath); err != nil {
			Log.Error().Err(err).Msg("Cannot remove shortcut")
		}
	}()

	// Disable auto updates
	utl.OverrideEnv("SLACK_NO_AUTO_UPDATES", "true")

	// Update deep link
	f, err := os.OpenFile(utl.PathJoin(electronBinPath, "resources", "app", "dist", "main.js"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Log.Error().Err(err).Msg("Cannot open main.js")
	}
	defer f.Close()
	if _, err := f.WriteString(` require('./portapps.js');`); err != nil {
		Log.Error().Err(err).Msg("Cannot write to main.js")
	}
	err = utl.WriteToFile(utl.PathJoin(electronBinPath, "resources", "app", "dist", "portapps.js"), `"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const path = require("path");
const { app } = require('electron');
app.on('browser-window-created', () => {
    if (app.listenerCount('browser-window-created') <= 10) {
        app.setAsDefaultProtocolClient('slack', process.execPath, ["--user-data-dir="+path.join(path.dirname(process.execPath), '..', '..', 'data')]);
    }
});`)
	if err != nil {
		Log.Error().Err(err).Msg("Cannot write to portapps.js")
	}

	app.Launch(os.Args[1:])
}
