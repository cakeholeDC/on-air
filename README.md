# ON AIR 🎙️🚨
![test-and-lint](https://github.com/cakeholeDC/on-air/actions/workflows/test-and-lint.yml/badge.svg)

ON AIR is an application for turning on a light to signal to my family that I am currently on a video call. 

This app deploys a script that runs as a cronjob and determines whether the indicator light should be ON or OFF, and then sends a signal to the controller for the device.

The applicaton can be configured to respond to two triggers: <!-- TODO: two => three -->
1. When a specified application(s) is running.
1. When Apple's `VDC_Assistant` (the little light next to the webcam) is activated. <!-- TODO: VDC_Assistant is external webcams only. Need to fix this. -->
<!-- TODO: implement audio trigger -->
<!-- 1. When Apples `IOAudioEngineState` is equal to 1 (true / enabled).  -->

This application utilizes _Apple HomeKit_ and _Apple Shortcuts_ to control the device.

<!-- TODO: If you're interested in using Smartthings to control your device, see the branch `smartthings` -->
<!-- TODO: If you're interested in using HomeAssistant to control your device, see the branch `hass` -->

## Pre-Requesites 
-  macOS Ventura 13.3+
- Apple [HomeKit](https://www.apple.com/home-app/)
- Apple [Shortcuts](https://support.apple.com/guide/shortcuts/welcome/ios)
- HomeKit compatible [device](https://www.apple.com/home-app/accessories/) (light, outlet, or switch recommended)
    - This device **must** already be paired with HomeKit

## Dependencies
- [homebrew](https://brew.sh/) => `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
- [pipx](https://pypa.github.io/pipx/) => `brew install pipx`
- [poetry](https://python-poetry.org/) => `pipx install poetry`
    - Note: Do not install poetry using homebrew
- [invoke](https://github.com/pyinvoke/invoke) => `pipx install invoke`
- python 3.12

## Setup
### Create Apple Shortcuts
Shortcuts can be created on macOS or iOS. 

Open the Shortcuts app and create two new shortcuts. One named "On Air" and one named "Off Air"
<!-- TODO: new shortcut for status, and poll that instead of a logfile. -->

1. For the action, select the **Home App**
1. From the list of actions, select **Control**
1. From the _Scenes and Accessories_ list, select your **device or scene**
1. Click **Next**
1. Select your **device's state**
1. Click **Done**
1. **Name the shortcut** and click **Done**

### Install
1. Install project dependencies and create .env
    - `inv install-dependencies` <!-- TODO: rename to setup project or something like that. -->
1. Write convenience scripts and bash aliases
    - `inv install`

### Configuration
The application is configured via the `.env` file.

| Variable     | Type        | Usage      |
| ------------ | ----------- | ---------- |
| TRIGGER_APPS | List ["str"]| [Process name(s)](#trigger-apps) to trigger the indicator light |
| USE_WEBCAM   | Boolean     | Enable trigger for webcam activation | 
<!-- TODO: implement audio trigger -->
<!-- | USE_AUDIO    | Boolean     | Enable trigger for microphone activation |  -->

#### Trigger Apps
1. Open the application(s) that you want to turn on the light.
1. Run the following command to find the application's process name. 
    - `inv discover-process-names -q {app-name}`
1. Find the process name in the output.

> **Note:** when searching for **_app-name_**, try using a short keyword like _"code"_ rather than _"Visual Studio Code"_
> 
> Sometimes, a process name is shortened to _"vscode"_ which would not show up with a multi word query.
>
> If you are unable to successfully query for the process, run the command without `-q {app-name}` to view *all* running processes. 

### Manual Usage
A cool party trick is being able to turn the light on or off from the command line. After running the install steps, you can manually control the light with the `onair` and `offair` commands. These commands do not override the cron process.

### Automation
To automate your ON AIR light, schedule a cron job to run the script `./run-app-status-light.sh`

#### Cron Scheduler (Recommended)
Use the builtin scheduler to write your crontab entry.
- `inv manage-cron [options]`

The scheduler accepts the options below:
```sh
Options:
  # the action to perform. Must be "add", "list" or "remove"
  -a STRING, --action=STRING
  # run every N minutes. Use 0 for every minute => only for action=add
  -i INT, --interval-min=INT
  # starting at hour (24H) => only for action=add
  -s INT, --start-hour=INT
  # ending at hour (24H) => only for action=add
  -e INT, --end-hour=INT
  # line number to remove => only for action=remove
  -l INT, --line-num=INT
```

Examples:
- `inv manage-cron --action add --interval-min 5 --start-hour 7 --end-hour 15`
- `inv manage-cron --action remove --line-num 3`
- `inv manage-cron --action list`

#### Manually (Legacy)
1. Open crontab list with `crontab -e`
2. Add your cron schedule to the crontab list
3. Save and Exit vim with `:wq`

Sample cron schedules have been provided in the in the `./cronjobs/` directory.
- Every 5th Minute, MON-FRI 9am-5pm
- Every Minute, Every Day

> Need help with cron scheduling? Check out [crontab.guru](https://crontab.guru/)


## Uninstall
1. Remove convenience scripts, bash aliases, and python venv
    - `inv uninstall`
1. Remove the crontab entry, if applicable
    - `inv manage-cron --action list`
    - `inv manage-cron --action remove --line-num {N}`
