<p align="center"><a href="https://github.com/crazy-max/slack-portable" target="_blank"><img width="100" src="https://github.com/crazy-max/slack-portable/blob/master/res/logo.png"></a></p>

<p align="center">
  <a href="https://github.com/crazy-max/slack-portable/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/slack-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/slack-portable/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/slack-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/crazy-max/slack-portable"><img src="https://img.shields.io/appveyor/ci/crazy-max/slack-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/slack-portable"><img src="https://goreportcard.com/badge/github.com/crazy-max/slack-portable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/slack-portable"><img src="https://img.shields.io/codacy/grade/8beee2b3463842f6ad27da362666e75c.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=4Y9DDQ8LDLMTL"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

A single EXE written in [Golang](https://golang.org/) to make [Slack](https://slack.com) portable on Windows systems.<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

## Installation

There are different kinds of artifacts :

* `slack-portable-{ia32,x64}-x.x.x-x-setup.exe` : Full portable release of Slack as a setup. **Recommended way**!
* `slack-portable-{ia32,x64}-x.x.x-x.7z` : Full portable release of Slack as a 7z archive.
* `slack-portable-{ia32,x64}.exe` : Only the portable binary (must be renamed `slack-portable.exe`)
* `SlackSetup-{ia32,x64}-x.x.x.exe` : The original setup from the [official website](https://slack.com/downloads/windows).
* `slack-{ia32,x64}-x.x.x-full.nupkg` : The original NUPKG file extracted from the original setup.

For a **fresh installation**, install `slack-portable-{ia32,x64}-x.x.x-x-setup.exe` where you want then run `slack-portable.exe`.

If **you have already installed Slack from the original setup**, do the same thing as a fresh installation and move the data located in `%APPDATA%\slack` to `data\AppData\Roaming\slack` then run `slack-portable.exe`.<br />
Then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) Slack from your computer.

**For an upgrade**, simply download and install the [latest setup](https://github.com/crazy-max/slack-portable/releases/latest).

## How can i help ?

We welcome all kinds of contributions :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
Any funds donated will be used to help further development on this project! :gift_heart:

[![Donate Paypal](https://raw.githubusercontent.com/crazy-max/slack-portable/master/res/paypal.png)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=4Y9DDQ8LDLMTL)

## License

MIT. See `LICENSE` for more details.<br />
USB icon credit to [Juliia Osadcha](https://www.iconfinder.com/Juliia_Os).
