# TODO

CORE
- Keygen: randomly generate a file.key in app_data. Default env var points here. Preserve ability to overwrite key with env var.
- Allow for a time interval. M-F between 9-17
    - config - add a config value for schedule, use CRONTAB syntax. Enforce in the binary.
    - How do I turn the light on/off outside of the scheduled window. Currently onair -o and onair -f will check the time to see if actions can be performed, and then not.
        - Should -o and -f override the cache? Shuold the cache only apply when run as `onair`? - what if we just removed the cache? That assumes false...

HASS
- Endpoint to list all devices by entity_id; CMD to access
- Are scenes/routines an option? It'd be cool to gradually change the dimmer of a smart light over 30-60 seconds.
- Rebuild Hass server.

MEDIA
- What if the light is turned on manually? How do we prevent the next scheduled check from turning it off?
    - Consider a cache/lock file to cache the light state.
    - How should this work?
- maybe: update media package c++ dependencies

CI/CD
- Do we need to run on MacOS anymore?

CONFIG
- Encrypt by default
- Currently, the app would require two config files to control two devices. Can we structure the yaml config in such a way that more than one device can be configured? Consider a structure as follows:

```yaml
# api
home_assistant_url: http://homeassistant.iot.home:8123
home_assistant_token: mYt0K3n

# triggers
onair_enable_camera: true
onair_enable_microphone: true

# devices
home_assistant_entities:
    - switch.onair
    - switch.lamp
    - switch.stage-light
```
this could allow for the option to invert the on/off state for a specific device (ie. turn one on and the other off)

```yaml
# devices
home_assistant_entities:
    - switch.onair:
    - switch.lamp:
        invert_state: true
    - switch.stage-light:
```

CLEANUP
- setup release actions
- improve go test coverage
