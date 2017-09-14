//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/app-portable.ico
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/logger"
)

const (
	NAME            = "slack-portable"
	APP_NAME        = "Slack"
	APP_DATA_FOLDER = "slack"
	APP_PROCESS     = "slack.exe"
)

func main() {
	// Current path
	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Error("Current path:", err)
	}

	// Logs folder
	logsPath := pathJoin(currentPath, "logs")
	logger.Info("Create logs folder", logsPath)
	err = os.Mkdir(logsPath, 777)
	if err != nil {
		logger.Error("Create logs folder:", err)
	}

	// Log file
	logfile, err := os.OpenFile(pathJoin(logsPath, NAME+".log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Error("Log file:", err)
	}

	// Init logger
	logger.Init(NAME, false, false, logfile)
	logger.Info("--------")
	logger.Infof("Starting %s...", NAME)
	logger.Infof("Current path: %s", currentPath)

	// Purge logs
	logsFolder, err := os.Open(logsPath)
	if err != nil {
		logger.Error("Open logs folder:", err)
	}
	defer logsFolder.Close()
	logsFiles, err := logsFolder.Readdir(-1)
	if err != nil {
		logger.Error("Read logs folder:", err)
	}
	logger.Infof("Reading %s...", logsPath)
	for _, logsFile := range logsFiles {
		if !strings.HasPrefix(logsFile.Name(), NAME) {
			os.Remove(pathJoin(logsPath, logsFile.Name()))
			logger.Infof("Deleted %s", pathJoin(logsPath, logsFile.Name()))
		}
	}

	// Find app folder
	logger.Infof("Lookup app folder in %s", currentPath)
	var appPath = ""
	rootFiles, _ := ioutil.ReadDir(currentPath)
	for _, f := range rootFiles {
		if strings.HasPrefix(f.Name(), "app-") && f.IsDir() {
			logger.Infof("App folder found: %s", f.Name())
			appPath = pathJoin(currentPath, f.Name())
			break
		}
	}
	if _, err := os.Stat(appPath); err == nil {
		logger.Infof("App path: %s", appPath)
	} else {
		logger.Error("App path does not exist")
	}

	// Init vars
	appExe := pathJoin(appPath, APP_PROCESS)
	dataPath := pathJoin(currentPath, "data")
	dataAppPath := pathJoin(dataPath, "AppData", "Roaming", APP_DATA_FOLDER)
	downloadsPath := pathJoin(currentPath, "downloads")
	slackSettingsPath := pathJoin(dataAppPath, "storage", "slack-settings")
	logger.Info("App executable:", appExe)
	logger.Info("Data path:", dataAppPath)
	logger.Info("Downloads path:", downloadsPath)
	logger.Info("Slack settings path:", slackSettingsPath)

	// Create data folder
	logger.Infof("Create data folder %s...", dataAppPath)
	err = os.MkdirAll(dataAppPath, 777)
	if err != nil {
		logger.Error("Create data folder:", err)
	}

	// Create download folder
	logger.Infof("Create download folder %s...", downloadsPath)
	err = os.MkdirAll(downloadsPath, 777)
	if err != nil {
		logger.Error("Create download folder:", err)
	}

	// Override USERPROFILE env var
	if err := os.Setenv("USERPROFILE", dataPath); err != nil {
		logger.Error("Cannot set USERPROFILE env var:", err)
	}

	// Change slack settings
	logger.Info("Update Slack settings...")
	if _, err := os.Stat(slackSettingsPath); err == nil {
		rawSettings, err := ioutil.ReadFile(slackSettingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			json.Unmarshal(rawSettings, &jsonMapSettings)
			logger.Info("Current settings:", jsonMapSettings)

			jsonMapSettings["resourcePath"] = pathJoin(appPath, "resources", "app.asar")
			jsonMapSettings["PrefSSBFileDownloadPath"] = downloadsPath
			jsonMapSettings["notificationMethod"] = "html"
			logger.Info("New settings:", jsonMapSettings)

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				logger.Error("Slack settings marshal:", err)
			}
			err = ioutil.WriteFile(slackSettingsPath, jsonSettings, 0644)
			if err != nil {
				logger.Error("Write Slack settings:", err)
			}
		}
	} else {
		logger.Errorf("Slack settings not found in %s", slackSettingsPath)
	}

	// Launch
	logger.Infof("Launch %s...", APP_NAME)
	execApp := exec.Command(appExe, "--log-file", "./")
	execApp.Dir = logsPath

	defer logfile.Close()
	execApp.Stdout = logfile
	execApp.Stderr = logfile

	if err := execApp.Start(); err != nil {
		logger.Error("Cmd Start:", err)
	}

	execApp.Wait()
}

func pathJoin(elem ...string) string {
	for i, e := range elem {
		if e != "" {
			return strings.Join(elem[i:], `\`)
		}
	}
	return ""
}
