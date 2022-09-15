#!/bin/bash

# gets UVCExtension.PowerLog output from the past 5 min
# returns VDCAssistant_Power_State
log show \
--predicate 'subsystem contains "com.apple.UVCExtension" and composedMessage contains "Post PowerLog"' \
--last 5m \
| grep "VDCAssistant_Power_State" \
| tail -1 \
| cut -d"=" -f2- \
| xargs
