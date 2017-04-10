//go:generate goversioninfo -icon=slack-portable.ico
package main

import (
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "os/exec"
  "path"
  "path/filepath"
  "strings"
  "syscall"

  "github.com/op/go-logging"
)

var log = logging.MustGetLogger("slack-portable")
var logFormat = logging.MustStringFormatter(`%{time} %{level:.4s} - %{message}`)

func main() {
  // Current path
  currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Error("Current path:", err)
  }

  // Log file
  logfile, err := os.OpenFile(path.Join(currentPath, "slack-portable.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
    log.Error("Log file:", err)
  }

  // Init logger
  logBackendStdout := logging.NewBackendFormatter(logging.NewLogBackend(os.Stdout, "", 0), logFormat)
  logBackendFile := logging.NewBackendFormatter(logging.NewLogBackend(logfile, "", 0), logFormat)
  logging.SetBackend(logBackendStdout, logBackendFile)
  log.Info("--------")
  log.Info("Starting slack-portable...")

  // Convert backslashes
  currentPath = path.Join(strings.Replace(string(currentPath), string(filepath.Separator), "/", -1))
  log.Info("Current path: " + currentPath)

  // Find app folder
  var appPath = ""
  rootFiles, _ := ioutil.ReadDir(currentPath)
  for _, f := range rootFiles {
    if (strings.HasPrefix(f.Name(), "app-") && f.IsDir()) {
      log.Info("App folder found:", f.Name())
      appPath = path.Join(currentPath, f.Name())
      break
    }
  }
  if _, err := os.Stat(appPath); err == nil {
    log.Info("App path:", appPath)
  } else {
    log.Warning("App path does not exist");
  }

  // Init vars
  var slackExe = path.Join(currentPath, "slack.exe")
  var dataPath = path.Join(currentPath, "data")
  var downloadsPath = path.Join(currentPath, "downloads")
  var symlinkPath = path.Clean(path.Join(os.Getenv("APPDATA"), "Slack"))
  //var symlinkPath = path.Join(currentPath, "data2")
  var slackSettingsPath = path.Join(dataPath, "storage", "slack-settings")
  log.Info("Slack executable:", slackExe)
  log.Info("Data path:", dataPath)
  log.Info("Downloads path:", downloadsPath)
  log.Info("Symlink path:", symlinkPath)
  log.Info("Slack settings path:", slackSettingsPath)

  // Check data folder
  if _, err := os.Stat(symlinkPath); err == nil {

    // Copy data folder
    if _, err := os.Stat(dataPath); os.IsNotExist(err) {
      os.Mkdir(dataPath, 777)
      err = copyDir(symlinkPath, dataPath)
      if err != nil {
        log.Error("Copying data folder:", err)
      }
    }

    // Rename old data folder if not a symlink
    fi, err := os.Lstat(symlinkPath)
    if err != nil {
      log.Error("Symlink lstat:", err)
    }
    if fi.Mode() & os.ModeSymlink != os.ModeSymlink {
      err = os.Chmod(symlinkPath, 0777)
      if err != nil {
        log.Error("Chmod symlink:", err)
      }
      err = os.Rename(symlinkPath, symlinkPath + "_old")
      if err != nil {
        log.Error("Renaming old data folder:", err)
      }
    }
  }

  // Create symlink
  log.Info("Creating symlink...")
  os.Remove(symlinkPath)
  cmd := exec.Command("cmd", "/c", "mklink", "/J", strings.Replace(symlinkPath, "/", "\\", -1), strings.Replace(dataPath, "/", "\\", -1))
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
  if err := cmd.Run(); err != nil {
    log.Error("Symlink:", err)
  }
  /*err = os.Symlink(dataPath, symlinkPath)
  if err != nil {
    log.Error(err)
  }*/

  // Check downloads folder
  log.Info("Checking downloads folder...")
  if _, err := os.Stat(downloadsPath); os.IsNotExist(err) {
    os.Mkdir(downloadsPath, 777)
  }

  // Change slack settings
  log.Info("Updating Slack settings...")
  if _, err := os.Stat(slackSettingsPath); err == nil {
    rawSettings, err := ioutil.ReadFile(slackSettingsPath)
    if err == nil {
      jsonMapSettings := make(map[string]interface{})
      json.Unmarshal(rawSettings, &jsonMapSettings)
      log.Info("Current settings:", jsonMapSettings)

      jsonMapSettings["resourcePath"] = strings.Replace(appPath + "/resources/app.asar", "/", "\\", -1)
      jsonMapSettings["PrefSSBFileDownloadPath"] = strings.Replace(downloadsPath, "/", "\\", -1)
      log.Info("New settings:", jsonMapSettings)

      jsonSettings, err := json.Marshal(jsonMapSettings)
      if err != nil {
        log.Error("Slack settings marshal:", err)
      }
      err = ioutil.WriteFile(slackSettingsPath, jsonSettings, 0644)
      if err != nil {
        log.Error("Write Slack settings:", err)
      }
    }
  } else {
    log.Warning("Slack settings not found in:", slackSettingsPath)
  }

  // Launch slack
  cmd = exec.Command(slackExe)
  if err := cmd.Run(); err != nil {
    log.Error("Launching:", err)
  }

  defer logfile.Close()
}

// src: https://gist.github.com/m4ng0squ4sh/92462b38df26839a3ca324697c8cba04
func copyFile(src, dst string) (err error) {
  in, err := os.Open(src)
  if err != nil {
    return
  }
  defer in.Close()

  out, err := os.Create(dst)
  if err != nil {
    return
  }
  defer func() {
    if e := out.Close(); e != nil {
      err = e
    }
  }()

  _, err = io.Copy(out, in)
  if err != nil {
    return
  }

  err = out.Sync()
  if err != nil {
    return
  }

  si, err := os.Stat(src)
  if err != nil {
    return
  }

  err = os.Chmod(dst, si.Mode())
  if err != nil {
    return
  }

  return
}

// src: https://gist.github.com/m4ng0squ4sh/92462b38df26839a3ca324697c8cba04
func copyDir(src string, dst string) (err error) {
  src = filepath.Clean(src)
  dst = filepath.Clean(dst)

  si, err := os.Stat(src)
  if err != nil {
    return err
  }
  if !si.IsDir() {
    return fmt.Errorf("source is not a directory")
  }

  _, err = os.Stat(dst)
  if err != nil && !os.IsNotExist(err) {
    return
  }

  err = os.MkdirAll(dst, si.Mode())
  if err != nil {
    return
  }

  entries, err := ioutil.ReadDir(src)
  if err != nil {
    return
  }

  for _, entry := range entries {
    srcPath := filepath.Join(src, entry.Name())
    dstPath := filepath.Join(dst, entry.Name())
    if entry.IsDir() {
      err = copyDir(srcPath, dstPath)
      if err != nil {
        return
      }
    } else {
      err = copyFile(srcPath, dstPath)
      if err != nil {
        return
      }
    }
  }

  return
}
