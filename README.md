# ON AIR
ON AIR is a python application for turning on a light when a specified application(s) is running. The purpose of this app is to turn a light on to signal to my family that I am currently on a video call.

This application utilizes the Samsung Smarttings API to control smart home devices utilizing the Smartthings ecosystem.

---
## Pre-Requesites
- [Samsumg Smartthings Hub](https://www.tomsguide.com/us/samsung-smart-things-v3,review-5809.html)
- Smartthings [Developer account](https://smartthings.developer.samsung.com/)
- Smartthings compatible [device](https://www.samsung.com/us/smart-home/compatible-devices/) (light bulb, outlet, led strip, etc.)
    - This device **must** already be configured within Smartthings.
- Mac OS 11+ (Untested on older OS)

## Dependencies
- python
- [poetry](https://python-poetry.org/)
- [python-smartthings](https://pypi.org/project/python-smartthings/)

---
## Setup
### Configure Environment
1. Create local environment
    - `poetry install`
2. Create your `.env` file
    - `cp .env.example .env`
3. Open `.env` in your preferred code editor.
4. Obtain a [smartthings personal access token](https://account.smartthings.com/tokens) fom the Smartthings developer hub.
5. Set your personal access token as the value of `TOKEN` in `.env`
6. Use the API to discover devices:
    - `poetry run python3 getDeviceList.py`
7. Capture your desired device's `Label` from the output and set it as the value of `DEVICE_NAME` in `.env`

### Configure Applications
1. Open the application(s) that you want to turn on the light.
2. Run the following command to find the application's proces name. 
    - Replace _{app-name}_ with the application name.
    - `poetry run python3 hostActions.py | grep {app-name}`
    > **Note:** when searching for **_app-name_**, try using a short keyword like _"code"_ rather than _"Visual Studio Code"_
    > 
    > Sometimes, a process name is shortened to _"vscode"_ which would not show up with a multi word query.
3. Find the process name in the output.
4. Copy & Paste the application name(s) into `VIDEO_LIST` in `.env`

### Testing the Configuration
Follow these steps to test your configiuration:
1. run `./run-on-air.sh` — the light should turn on.
2. run `./run-off-air.sh` — the light should turn off.
3. Open the trigger application(s) 
    - run `./run-app-status-light.sh` — the light should turn on. 
4. Quit the trigger application(s)
    - run `./run-app-status-light.sh` — the light should turn off.

---
## Automation
To automate your ON AIR light, schedule a cron job to run the script `./run-app-status-light.sh`

Sample cron schedules have been provided in the in the `./cronjobs/` directory.
- Every 5th Minute, MON-FRI 9am-5pm
- Every Minute, Every Day


### Adding a cron job
1. Open crontab list with `crontab -e`
2. Add your cron schedule to the crontab list
3. Save and Exit vim with `:wq`

> Need help with cron scheduling? Check out [crontab.guru](https://crontab.guru/)


---
## Tips & Tricks
### Terminal Aliases
A cool party trick is being able to turn the light on or off from the command line. This can be accomplished by adding aliases to your `$PATH` that run the on/off scripts.

Run `./scripts/deploy_terminal_alias` to quickly configure your `$PATH` with `onair` and `offair` aliases for controlling the light. 
