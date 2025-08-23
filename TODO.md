# TODO

CORE
- ✅ Hook up run logic for should light be on. `onair` => default cfg => do the thing.
- ✅ Constants
- Keygen: randomly generate a file.key in app_data. Default env var points here. Preserve ability to overwrite key with env var.

UX
- What does a new user do first? WIll they need a readme, or will the binary have a good enough ux?

HASS
- ✅ Wire up CMD on/off to API on/off
- Endpoint to list all devices by entity_id; CMD to access
- Are scenes/routines an option? It'd be cool to gradually change the dimmer of a smart light over 30-60 seconds.
- Rebuild Hass server.

CLEANUP
- remove python code entirely
- setup release actions
- improve go test coverage

MEDIA
- What if the light is turned on manually? How do we prevent the next cron job from turning it off?
Consider a cache file to cache the light state. How should this work?

- maybe: update media package c++ dependencies

CONFIG
- ✅ Currently, we auto create the config if it's not there. We should probably fail and offer a config --create option. 

With this change, we can then invoke the binary `onair` with no sub commands or flags, to run the core functionality of: "should the light be on/off? Ok, let's turn it on/off". 

That way, if there's no config, running `onair` will fail with a message that **"configuration is required for this application. please `run onair config --create`"** to create one.

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