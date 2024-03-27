#!/bin/bash

# log show  --last 1m | grep -i "AppleH13CamIn::setPowerStateGated"
# log show  --predicate 'composedMessage contains "setPowerStateGated"' --last 4m


#* ! SUCCESS !
#  ~ kcole $ log show  --predicate 'composedMessage contains "setPowerStateGated"' --last 1m
# Filtering the log data using "composedMessage CONTAINS "setPowerStateGated""
# Skipping info and debug messages, pass --info and/or --debug to include.
# Timestamp                       Thread     Type        Activity             PID    TTL  
# 2024-03-26 10:31:41.803180-0700 0x4a46a    Default     0x0                  0      0    kernel: (AppleH11ANEInterface) ANE0: setPowerStateGatedPriv :H11ANEIn::setPowerStateGated: 1
# 2024-03-26 10:31:41.851301-0700 0x4a46a    Default     0x0                  0      0    kernel: (AppleH11ANEInterface) ANE0: setPowerStateGatedPriv :H11ANEIn::setPowerStateGated: 0
# 2024-03-26 10:31:42.865466-0700 0x4a464    Default     0x0                  0      0    kernel: (AppleH11ANEInterface) ANE0: setPowerStateGatedPriv :H11ANEIn::setPowerStateGated: 1


#? AUDIO??
#  ~ kcole $ ioreg -l |grep IOAudioEngineState
#     | |   | | |   | |   |     |   "IOAudioEngineState" = 1

# NOTE: this keeps MEET open, but otherwise the microphone must be ACTIVE.
# IE. photobooth does not make this true. Photobooth, use video.
#*!! ioreg -l | grep IOAudioEngineState | cut -d "=" -f2- | xargs


log show --predicate 'subsystem contains "AppleH13CamIn"' --last 1m

# log show \
# --predicate 'subsystem contains "AppleH13CamIn"' \
# --last 5m \
# | grep "VDCAssistant_Power_State" \
# | tail -1 \
# | cut -d"=" -f2- \
# | xargs
