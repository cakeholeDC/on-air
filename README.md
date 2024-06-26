# ON AIR 🎙️🚨
![test-and-lint](https://github.com/cakeholeDC/on-air/actions/workflows/test-and-lint.yml/badge.svg)

ON AIR is an application for turning on a light to signal to my family that I am currently on a video call. 

This app deploys a cronjob script that determines whether the light should be ON or OFF, and then sends a signal to the integrated IoT system.

The application can be configured to respond to three triggers:
1. When a specified application(s) is running.
1. When the webcam is active. Webcam activity is determined by:
    - Apple's `VDC_Assistant` (a USB webcam) is activated.
    - Apple's `AppleH13CamIn::setGetPowerStateGated` (the onboard camera) is changed.
1. When Apples `IOAudioEngineState` is equal to 1 (true / enabled). 

<!-- TODO: Smartthings: see the branch `smartthings` -->
This application supports integrations with _Home Assistant_, and _Apple HomeKit_.

## Pre-Requisites  
Each supported integration has it's own requirements listed below.

### Home Assistant
-  macOS Monterey 12.7+
- [Home Assistant](https://www.home-assistant.io/)
- Home Assistant compatible [device](https://www.home-assistant.io/integrations/) (light, outlet, or switch recommended)
    - This device **must** already be paired with Home Assistant

### Apple HomeKit
> **Note:** This app only supports the [new HomeKit architecture](https://support.apple.com/en-us/102287) released in 2023. [Read More](https://www.reddit.com/r/HomeKit/comments/zsir3n/explanation_on_new_homekit_architecture/)
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
- [python 3.11](https://www.python.org/downloads/release/python-3110/) => `brew install python@3.11`

## Setup
### Home Assistant: Setup Entity
1. Add your Server URL to `.env` as `HASS_SERVER_URL`.
1. Add the necessary Home Assistant [integration](https://www.home-assistant.io/getting-started/integration/) for your device.
    - Capture the [`entity_id`](https://www.home-assistant.io/docs/configuration/customizing-devices/) you wish to control
    - Add the `entity_id` to `.env` as `HASS_ENTITY_ID`.
1. Create a Home Assistant [Long Lived Access Token](https://developers.home-assistant.io/docs/auth_api/#long-lived-access-token)
    - Add the `token` to `.env` as `HASS_ENTITY_ID`.

### HomeKit: Create Apple Shortcuts
Shortcuts can be created on macOS or iOS. 

Open the Shortcuts app and create two new shortcuts. One to turn your device ON, and one to turn it OFF. 

1. For the action, select the **Home App**
1. From the list of actions, select **Control**
1. From the _Scenes and Accessories_ list, select your **device or scene**
1. Click **Next**
1. Select your **device's state**
1. Click **Done**
1. **Name the shortcut** and click **Done**

Add the names of these shortcuts to `.env` as `SHORTCUT_ON` and `SHORTCUT_OFF`.

### Install
1. Install project dependencies and create .env
    - `inv project-setup`
1. Write convenience scripts and bash aliases
    - `inv app-install`
1. Setup the cronjob (if desired)
    - [#cron-scheduler](#cron-scheduler-recommended)

### Configuration
The application is configured via the `.env` file.

| Variable       | Type        | Usage      |
| -------------- | ----------- | ---------- |
| TRIGGER_APPS   | List ["str"]| [Process name(s)](#trigger-apps) to trigger the device |
| ENABLE_VIDEO   | Boolean     | Enable trigger for webcam activation | 
| ENABLE_AUDIO   | Boolean     | Enable trigger for microphone activation | 
| SMART_HOME_TYPE   | String     | The IoT system to integrate | 
| DEVICE_CACHE   | String      | device cache filename |
| VIDEO_CACHE    | String      | video cache filename |
| HASS_SERVER_URL    | String      | http://hass.server:8123 |
| HASS_API_TOKEN    | String      | Home Assistant [Long Lived Access Token](https://developers.home-assistant.io/docs/auth_api/#long-lived-access-token) |
| HASS_ENTITY_ID    | String      | switch.identifier |
| SHORTCUT_ON    | String      | name of Homekit shortut for device on |
| SHORTCUT_OFF   | String      | name of Homekit shortut for device off |

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
1. Remove convenience scripts and bash aliases
    - `inv app-uninstall`
1. Remove the crontab entry, if applicable
    - `inv manage-cron --action list`
    - `inv manage-cron --action remove --line-num {N}`
