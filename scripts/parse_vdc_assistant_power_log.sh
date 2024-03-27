#!/bin/bash

# TODO: find the same thing for audio.

# TODO: look for something like: AppleH13CameraInterface in the logs
# $ log show  --last 1m | grep -i "AppleH13CamIn::setPowerStateGated" 

#  ~ kcole $ log show  --last 1m | grep -i "AppleH13CamIn::setPowerStateGated"
# 2024-03-26 07:57:35.196882-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 1
# 2024-03-26 07:57:35.196892-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - Transitioned fPowerEvent from 1 to onPending
# 2024-03-26 07:57:36.236747-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - change fPowerEvent from 1 to None
#  ~ kcole $ log show  --last 1m | grep -i "AppleH13CamIn::setPowerStateGated" 
# 2024-03-26 07:57:35.196882-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 1
# 2024-03-26 07:57:35.196892-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - Transitioned fPowerEvent from 1 to onPending
# 2024-03-26 07:57:36.236747-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - change fPowerEvent from 1 to None
# 2024-03-26 07:57:40.241043-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 0
# 2024-03-26 07:57:40.241048-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn off) - change fPowerEvent from 2 to PowerOffPending, HasActiveKernelClients=0
# 2024-03-26 07:57:40.292833-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn off) - change fPowerEvent from 2 to None
#  ~ kcole $ log show  --last 1m | grep -i "AppleH13CamIn::setPowerStateGated" 
# 2024-03-26 07:57:35.196882-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 1
# 2024-03-26 07:57:35.196892-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - Transitioned fPowerEvent from 1 to onPending
# 2024-03-26 07:57:36.236747-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - change fPowerEvent from 1 to None
# 2024-03-26 07:57:40.241043-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 0
# 2024-03-26 07:57:40.241048-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn off) - change fPowerEvent from 2 to PowerOffPending, HasActiveKernelClients=0
# 2024-03-26 07:57:40.292833-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn off) - change fPowerEvent from 2 to None
# 2024-03-26 07:57:46.389265-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 1
# 2024-03-26 07:57:46.389271-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - Transitioned fPowerEvent from 1 to onPending
# 2024-03-26 07:57:47.440635-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - change fPowerEvent from 1 to None
#  ~ kcole $ log show  --last 1m | grep -i "AppleH13CamIn::setPowerStateGated" 
# 2024-03-26 07:57:35.196882-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 1
# 2024-03-26 07:57:35.196892-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - Transitioned fPowerEvent from 1 to onPending
# 2024-03-26 07:57:36.236747-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - change fPowerEvent from 1 to None
# 2024-03-26 07:57:40.241043-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 0
# 2024-03-26 07:57:40.241048-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn off) - change fPowerEvent from 2 to PowerOffPending, HasActiveKernelClients=0
# 2024-03-26 07:57:40.292833-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn off) - change fPowerEvent from 2 to None
# 2024-03-26 07:57:46.389265-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated - 1
# 2024-03-26 07:57:46.389271-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - Transitioned fPowerEvent from 1 to onPending
# 2024-03-26 07:57:47.440635-0700 0x36e15    Default     0x0                  0      0    kernel: (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - change fPowerEvent from 1 to None

# gets UVCExtension.PowerLog output from the past 5 min
# returns VDCAssistant_Power_State
log show \
--predicate 'subsystem contains "com.apple.UVCExtension" and composedMessage contains "Post PowerLog"' \
--last 5m \
| grep "VDCAssistant_Power_State" \
| tail -1 \
| cut -d"=" -f2- \
| xargs
