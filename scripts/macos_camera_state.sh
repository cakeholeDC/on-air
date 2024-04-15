#!/bin/bash

# This script returns the status of the Webcam by:
#   1. Parsing the system logs for 'VDCAssistant_Power_State' entries within the past 5 minutes.
#     - The VDCAssistant is what controls the power state for external webcams. 
#     - Read More: https://keith.github.io/xcode-man-pages/VDCAssistant.8.html
#     - This state does not reflect the state of the builtin camera.
#     - returns 'on' or 'off'
#   2. Parsing the system logs for 'AppleH13CamIn::setGetPowerStateGated' entries
#     - This is the command to set the builtin camera interface's power state
#     - It reflects that the webcam has been turned on or off
#     - returns 'on' or 'off'

# If either camera is 'on', this script returns "CAMERA=on"

# Example:
# ---- USB CAMERA
# $ log show --last 5m | grep "VDCAssistant"
# > "VDCAssistant_Power_State" = On;
# ---- BUILTIN CAMERA
# $ log show --last 5m | grep "AppleH13CamIn::setPowerStateGated"
# > : (AppleH13CameraInterface) AppleH13CamIn::setPowerStateGated (turn on) - Transitioned fPowerEvent from 1 to onPending

# TODO: optionally pass this value into the script
N_MINUTES=5

USB=$(log show \
  --predicate 'subsystem contains "com.apple.UVCExtension" and composedMessage contains "Post PowerLog"' \
  --last "${N_MINUTES}m" \
  | grep "VDCAssistant_Power_State" \
  | tail -1 \
  | cut -d"=" -f2- \
  | xargs \
  | sed "s/;//" \
  | awk '{print tolower($0)}'
)
ONBOARD=$(log show \
  --predicate 'process == "kernel" && (eventMessage contains "AppleH13CamIn::setPowerStateGated")' \
  --last "${N_MINUTES}m" \
  | grep "AppleH13CamIn::setPowerStateGated" \
  | tail -1 \
  | awk -F"AppleH13CamIn::setPowerStateGated" '{print $2}' \
  | cut -d"(" -f2 | cut -d")" -f1 \
  | sed "s/turn//" \
  | xargs
)

# # Uncomment for the status of each device independently
# echo "USB=$USB"
# echo "ONBOARD=$ONBOARD"


if [[ ! "$USB" ]]
then
  USB="null"
fi

if [[ ! "$ONBOARD" ]]
then
  ONBOARD="null"
fi


if [[ $USB == "on" || $ONBOARD == "on" ]]; then
  # echo "CAMERA=on"
  echo "True"
else
  if [[ $USB == "null" && $ONBOARD == "null" ]]; then
    # no activitiy in the last n minutes.
    echo "check_cache"
  else
    # echo "CAMERA=off"
    echo "False"
  fi
fi
