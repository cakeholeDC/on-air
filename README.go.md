# ON AIR 🎙️🚨
![test-and-lint](https://github.com/cakeholeDC/on-air/actions/workflows/test-and-lint.yml/badge.svg)

ON AIR is an application for interacting with IoT devices. Along with manual control, this application can detect whether the systems camera or microphone are enabled and turn on an indicator light. 

This app can also deploy a cronjob to run on an interval.

# Services

## Config
Onair uses a configuration file. Configuration files are 1:1 with IoT devices.

The configuration file holds values for interfacing with IoT integrations such as Home Assistant, Smartthings, HomeKit.

The default path for the config file is `$HOME/.config/onair/onair.cfg`

To specify a configuration file, set the environment variable `ONAIR_CONFIG_FILE_PATH` at runtime.

You can have multiple config files. Why would you want more than one config file? Let's say that you want two devices to turn on under different circumstances? Each device would require it's own config file. 

## Encryption
Configuration files can be encrypted to protect secrets. Set the encryption key with the environment variable `ONAIR_CONFIG_ENCRYPTION_KEY` - this is a string. If an encryption key is not provided, configuration files will not be encrypted.

The key should be an AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.

Keygen:
```bash
BYTES=32
ONAIR_CONFIG_ENCRYPTION_KEY=$(openssl rand -hex $BYTES)
```

## Logging
onair logs to `$HOME/.config/onair/onair.cfg`.

The log location can be changed with the `ONAIR_LOG_FILEPATH` env var.


| service | env var | type | default |
|---------|---------|------|---------|
| config  | `ONAIR_CONFIG_FILE_PATH` | string(Path) | `$HOME/.config/onair/onair.cfg` |
| config  | `ONAIR_CONFIG_ENCRYPTION_KEY` | string(16, 24, or 32 bytes) | null |
| logger  | `ONAIR_LOG_FILEPATH` | string(path) |  `$HOME/.config/onair/onair.log` |

Run `onair config` to see the config options.

## Home Assistant

## Service

## Service

## Service

## Launch Daemon

Schedule onair to run using launch daemon.

The file `local.onair.plist` contains an interval in seconds:
```
    <key>StartInterval</key>
    <integer>30</integer>
    <key>RunAtLoad</key>
```

Set this value to your desired interval.

Then run `make launchd-load` to load the service.

If you want to load the service at the top of the next minute to ensure your interval aligns with the minute, run `make launchd-interval`

Run `make launchd-unload` to remove the launch daemon.

## Run Binary as sudo without password:
```sh
sudo visudo
```

Find the section for "root and users in group wheel can run anything on any machine as any user" and add the following below the line: `%admin          ALL = (ALL) ALL`

```txt
%admin          ALL = (ALL) NOPASSWD: /usr/local/bin/onair
```

This will allow all admins to execute the binary without a password.