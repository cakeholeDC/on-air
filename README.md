# ON AIR 🎙️🚨
ON AIR is a python application for turning on a light to signal to my family that I am currently on a video call. 

The applicaton is configured to respond when a specified application(s) is running. It can also be configured to respond when Apple's `VDC_Assistant` (the little light next to the webcam) is activated.

This application utilizes _Apple HomeKit_ and _Apple Shortcuts_ to control the device.

![test-and-lint](https://github.com/cakeholeDC/on-air/actions/workflows/test-and-lint.yml/badge.svg)
<!-- ######### TODO: refactor to use Homekit and shortcuts -->
<!-- shortcuts run "On Air" -->


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

### Configure Environment
<!-- TODO: update with invoke -->
1. Create local environment
    - `inv install-dependencies`
1. Create your `.env` file
    - `cp .env.example .env`
1. Open `.env` in your preferred code editor.
1. run `inv deploy-scripts` to write the scripts to your home dir
1. run `inv TODO` to create your bash aliases.

### Configure Applications
1. Open the application(s) that you want to turn on the light.
1. Run the following command to find the application's proces name. 
    - Replace _{app-name}_ with the application name.
    - `inv discover-process-names -q {app-name}`
    > **Note:** when searching for **_app-name_**, try using a short keyword like _"code"_ rather than _"Visual Studio Code"_
    > 
    > Sometimes, a process name is shortened to _"vscode"_ which would not show up with a multi word query.
    >
    > If you still cannot find your process name, run this command without adding ` | grep {app-name}` to see all processes. 
1. Find the process name in the output.
1. Copy & Paste the application name(s) into `TRIGGER_APPS` in `.env`

<!-- ### Testing the Configuration
<!-- TODO: update -->
<!-- Follow these steps to test your configiuration:
1. run `./run-on-air.sh` — the light should turn on.
2. run `./run-off-air.sh` — the light should turn off.
3. Open the trigger application(s) 
    - run `./run-app-status-light.sh` — the light should turn on. 
4. Quit the trigger application(s)
    - run `./run-app-status-light.sh` — the light should turn off. -->

---
## Automation
To automate your ON AIR light, schedule a cron job to run the script `./run-app-status-light.sh`

Sample cron schedules have been provided in the in the `./cronjobs/` directory.
- Every 5th Minute, MON-FRI 9am-5pm
- Every Minute, Every Day


### Adding a cron job
#### Scheduler (Recommended)
Use the builtin scheduler to write your crontab entry:

Run `inv manage-cron` to setup your crontab entry. The scheduler accepts the options below:
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

> Need help with cron scheduling? Check out [crontab.guru](https://crontab.guru/)

---
## Tips & Tricks
### Terminal Aliases
<!-- TODO: update with invoke -->
A cool party trick is being able to turn the light on or off from the command line. This can be accomplished by adding aliases to your `$PATH` that run the on/off scripts.

<!-- TODO: make this invoke -->
Run `./scripts/deploy_terminal_alias` to quickly configure your `$PATH` with `onair` and `offair` aliases for controlling the light. 
