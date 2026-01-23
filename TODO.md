# TODO

CORE
- Keygen: randomly generate a file.key in app_data. Default env var points here. Preserve ability to overwrite key with env var.

HASS
- Endpoint to list all devices by entity_id; CMD to access
- Are scenes/routines an option? It'd be cool to gradually change the dimmer of a smart light over 30-60 seconds.
- Rebuild Hass server.

MEDIA
- Consider a lock file to turn the light on for N time.
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
