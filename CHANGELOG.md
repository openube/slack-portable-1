# Changelog

## 3.4.0-19 (2019/04/19)

* Slack 3.4.0
* Portapps 1.22.1

## 3.3.8-18 (2019/03/19)

* Create data folder at first launch
* Portapps 1.20.3

## 3.3.8-17 (2019/03/09)

* Slack 3.3.8
* Portapps 1.20.2

## 3.3.7-41 (2019/01/18)

* Slack 3.3.7
* [New Slack logo](https://slackhq.com/say-hello-new-logo)

## 3.3.3-40 (2018/10/24)

* Fix Windows desktop notifications

## 3.3.3-39 (2018/10/02)

* Slack 3.3.3
* Go 1.11
* Use [go mod](https://golang.org/cmd/go/#hdr-Module_maintenance) instead of `dep`

## 3.3.1-38 (2018/09/01)

* Slack 3.3.1

## 3.3.0-37 (2018/08/15)

* Slack 3.3.0

## 3.2.0-36 (2018/06/06)

* Slack 3.2.0

## 3.1.1-35 (2018/04/05)

* Slack 3.1.1

## 3.1.0-34 (2018/03/10)

* Slack 3.1.0

## 3.0.5-33 (2018/02/11)

* Move ia32/x64 to win32/win64 for arch def
* Remove nupkg file
* Add portapp.json file
* Uncheck run app in setup

## 3.0.5-32 (2018/02/09)

* Ability to pass custom args to the portable process

## 3.0.5-31 (2018/01/21)

* Slack 3.0.5

## 3.0.3-30 (2018/01/13)

* Rebuild electron.asar to push data directly in `data` subfolder
* No need to override USERPROFILE environment variable anymore

> :warning: **UPGRADE NOTES**
> * Move everything in `data\AppData\Roaming\slack\*` to `data` folder and remove folder `data\AppData`.

## 3.0.3-29 (2018/01/12)

* Slack 3.0.3

## 3.0.0-28 (2017/12/09)

* Slack 3.0.0

## 2.9.0-27 (2017/11/23)

* Move app to a subfolder
* Reduce dependencies to avoid heuristic detection
* Add UPX compression

> :warning: **UPGRADE NOTES**
> * Delete all files and folders in root folder except `data` folder.

## 2.9.0-26 (2017/11/19)

* Slack 2.9.0
* Move repository to [Portapps](https://github.com/portapps) org
* Switch to [Golang Dep](https://github.com/golang/dep) as dependency manager
* Go 1.9.1

## 2.8.2-25 (2017/10/19)

* Slack 2.8.2

## 2.8.1-24 (2017/10/01)

* Slack 2.8.1

## 2.8.0-23 (2017/09/14)

* Slack 2.8.0
* New instance opened with Windows notification method (Issue #8)

## 2.7.1-22 (2017/09/05)

* New logger
* Override USERPROFILE env var instead of using symlink to APPDATA to store data
* Do not migrate old data folder from APPDATA
* Reduce dependencies and system calls to avoid heuristic detection

> :warning: **UPGRADE NOTES**
> * Move the content of `data\*` in `data\AppData\Roaming\slack\`
> * Remove symlink `%APPDATA%\slack`

## 2.7.1-21 (2017/08/26)

* Go 1.9
* Add support guidelines

## 2.7.1-19 (2017/08/16)

* Slack 2.7.1

## 2.7.0-18 (2017/08/09)

* Slack 2.7.0
* ia32 releases now included
* Small refactoring

## 2.6.5-16 (2017/07/13)

* Slack 2.6.5

## 2.6.3-15 (2017/06/29)

* Slack 2.6.3
* Admin privileges not required for Setup

## 2.6.2-14 (2017/05/24)

* Slack 2.6.2
* Download Glide as an external lib

## 2.6.0-beta189985592-13 (2017/05/14)

* Provide the nupkg file in the release
* Add Go report card 
* Add Glide dependency manager for Go
* Standard code organization

## 2.6.0-beta189985592-12 (2017/05/03)

* Error on first install with InnoSetup (Issue #7)

## 2.6.0-beta189985592-10 (2017/04/28)

* Create a setup for portable installation (Issue #6)

## 2.6.0-beta189985592-9 (2017/04/28)

* Add Slack logging (Issue #4)
* Add ability to select the Slack version (Issue #3)
* Slack crashes because of Update.exe missing (Issue #2)

## 2.6.0-beta189985592-7 (2017/04/27)

* Not launched if data folder does not exists (Issue #1)
* Improve logging
* Slack 2.6.0-beta189985592

## 2.5.2-5 (2017/04/10)

* GetFileAttributesEx error when no data folder exists

## 2.5.2-1 (2017/04/09)

* Initial version
