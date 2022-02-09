import os
from LightControl import loop, lightsOff, lights_object, device_name
from hostActions import write_log_file

loop.run_until_complete( lightsOff(lights_object, device_name) )

write_log_file(False)
