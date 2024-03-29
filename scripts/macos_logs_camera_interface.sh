#!/bin/bash

# log show  --last 1m | grep -i "AppleH13CamIn::setPowerStateGated"
# log show  --predicate 'composedMessage contains "setPowerStateGated"' --last 4m




#? AUDIO - NEEDS TO BE PARSED.
#  ~ kcole $ ioreg -l |grep IOAudioEngineState
#     | |   | | |   | |   |     |   "IOAudioEngineState" = 1

# NOTE: this keeps MEET open, but otherwise the microphone must be ACTIVE.
# IE. photobooth does not make this true. Photobooth, use video.
#* AUDO SUCCESS !
# NOTE: this isn't log parsing, it's a device status. That's a good thing.
# SOURCE: https://stackoverflow.com/a/66070592
AUDIO=$(ioreg -l | grep IOAudioEngineState | cut -d "=" -f2- | xargs)

#* Video NOTES
#  ~ kcole $ log show  --predicate 'composedMessage contains "setPowerStateGated"' --last 1m
# Filtering the log data using "composedMessage CONTAINS "setPowerStateGated""
# Skipping info and debug messages, pass --info and/or --debug to include.
# Timestamp                       Thread     Type        Activity             PID    TTL  
# 2024-03-26 10:31:41.803180-0700 0x4a46a    Default     0x0                  0      0    kernel: (AppleH11ANEInterface) ANE0: setPowerStateGatedPriv :H11ANEIn::setPowerStateGated: 1
# 2024-03-26 10:31:41.851301-0700 0x4a46a    Default     0x0                  0      0    kernel: (AppleH11ANEInterface) ANE0: setPowerStateGatedPriv :H11ANEIn::setPowerStateGated: 0
# 2024-03-26 10:31:42.865466-0700 0x4a464    Default     0x0                  0      0    kernel: (AppleH11ANEInterface) ANE0: setPowerStateGatedPriv :H11ANEIn::setPowerStateGated: 1

# NO RESULTS
#log show --predicate 'subsystem contains "AppleH13CamIn"' --last 1m
#* VIDEO - NEEDS TO BE PARSED:
# log show  --predicate 'composedMessage contains "setPowerStateGated"' --last 1m

#* VIDEO SUCCESS!!!!
# ! WARNING: this seems inconsistent. Sometimes it returns 0 while the camera is active.
VIDEO=$(log show  --predicate 'composedMessage contains "setPowerStateGated"' --last 5m | grep "H11ANEIn::setPowerStateGated" | tail -1 | awk -F'setPowerStateGated:' '{print $2}' | xargs)

# POSSIBLE BETTER VIDEO:
# SOURECE: https://stackoverflow.com/a/69736932
#  - https://stackoverflow.com/a/74870752
# kcole $ log show --predicate 'sender contains "appleh13camerad" and (composedMessage contains "PowerOnCamera" or composedMessage contains "PowerOffCamera")' --last 5m
# Filtering the log data using "sender CONTAINS "appleh13camerad" AND (composedMessage CONTAINS "PowerOnCamera" OR composedMessage CONTAINS "PowerOffCamera")"
# Skipping info and debug messages, pass --info and/or --debug to include.
# Timestamp                       Thread     Type        Activity             PID    TTL  
# 2024-03-28 15:53:07.216267-0700 0x1272b7   Default     0x0                  554    0    appleh13camerad: PowerOnCamera : ISP_PowerOnCamera: powered on camera
# 2024-03-28 15:53:11.224278-0700 0x1311b9   Default     0x0                  554    0    appleh13camerad: PowerOffCamera : ISP_PowerOffCamera : 0x00000000
# 2024-03-28 15:53:20.266091-0700 0x1311b9   Default     0x0                  554    0    appleh13camerad: PowerOnCamera : ISP_PowerOnCamera: powered on camera
# 2024-03-28 15:53:20.903273-0700 0x134553   Default     0x0                  554    0    appleh13camerad: PowerOnCamera : Camera is already ON, returning
# 2024-03-28 15:53:20.977191-0700 0x1272b7   Default     0x0                  554    0    appleh13camerad: PowerOnCamera : Camera is already ON, returning
# 2024-03-28 15:53:32.269517-0700 0x1272b7   Default     0x0                  554    0    appleh13camerad: PowerOffCamera : ISP_PowerOffCamera : 0x00000000

###! OR
# https://stackoverflow.com/a/77920548
# kcole $ log show --predicate 'process == "kernel" && (eventMessage contains "AppleH13CamIn::power_off_hardware" || eventMessage contains "AppleH13CamIn::power_on_hardware")' --last 5m
# Filtering the log data using "process == "kernel" AND (composedMessage CONTAINS "AppleH13CamIn::power_off_hardware" OR composedMessage CONTAINS "AppleH13CamIn::power_on_hardware")"
# Skipping info and debug messages, pass --info and/or --debug to include.
# Timestamp                       Thread     Type        Activity             PID    TTL  
# 2024-03-28 15:53:19.220771-0700 0x1008c6   Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::power_on_hardware
# 2024-03-28 15:53:19.229820-0700 0x1008c6   Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::power_on_hardware - ISP-ANE networks, no DART remapping for ISP side during poweron
# 2024-03-28 15:53:32.269686-0700 0x1008c6   Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::power_off_hardware
# 2024-03-28 15:54:20.212656-0700 0x1008c6   Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::power_on_hardware
# 2024-03-28 15:54:20.221958-0700 0x1008c6   Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::power_on_hardware - ISP-ANE networks, no DART remapping for ISP side during poweron

echo "AUDIO=$AUDIO"
echo "VIDEO=$VIDEO"

# log show \
# --predicate 'subsystem contains "AppleH13CamIn"' \
# --last 5m \
# | grep "VDCAssistant_Power_State" \
# | tail -1 \
# | cut -d"=" -f2- \
# | xargs
