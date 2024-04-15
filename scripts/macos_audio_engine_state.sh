#!/bin/bash

# NOTE: unlike the camera_state script, this is not parsing system logs.
#       this script reads the raw IO device state.

# This script returns the status of the Microphone by:
#   1. Parsing the io registry (ioreg) for 'IOAudioEngineState'
#     - This is the raw audio IO device's state
#     - Read More: https://stackoverflow.com/a/66070592
#     - Returns 0 (false) or 1 (true)

# If the AudioEngine is active, this script returns "AUDIO=on"

# Example:
# $ ioreg -l | grep IOAudioEngineState
# > "IOAudioEngineState" = 1

# parse the response to extract the IO value
AUDIO_STATE=$(/usr/sbin/ioreg -l | grep IOAudioEngineState | cut -d "=" -f2- | xargs)
if [ 1 == "$AUDIO_STATE" ]; then
    # echo "AUDIO=on"
    echo "True"
else 
    # echo "AUDIO=off"
    echo "False"
fi
